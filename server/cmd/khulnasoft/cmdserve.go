// Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/khulnasoft/khulnasoft/server/config"
	"github.com/khulnasoft/khulnasoft/server/locales"
	"github.com/khulnasoft/khulnasoft/server/persistence"
	"github.com/khulnasoft/khulnasoft/server/persistence/relational"
	"github.com/khulnasoft/khulnasoft/server/public"
	"github.com/khulnasoft/khulnasoft/server/router"
	"golang.org/x/crypto/acme/autocert"
)

var serveUsage = `
"serve" starts the KhulnaSoft instance and listens to the configured port(s).
Configuration is sourced either from the envfile given to -envfile or a file
called khulnasoft.env in the default lookup hierarchy (this applies to Linux and
Darwin only):

- In the current working directory
- In ~/.config
- In $XDG_CONFIG_HOME
- In /etc/khulnasoft

In case no envfile is found or given, the environment variables already set are
used. More documentation about configuration KhulnaSoft can be found at:
https://docs.khulnasoft.com/running-khulnasoft/configuring-the-application/

To find out about the configuration that would currently by applied, use
the "khulnasoft debug" subcommand.

Usage of "serve":
`

func cmdServe(subcommand string, flags []string) {
	cmd := flag.NewFlagSet(subcommand, flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), serveUsage)
		cmd.PrintDefaults()
	}
	var (
		envFile = cmd.String("envfile", "", "the env file to use")
	)
	cmd.Parse(flags)
	a := newApp(false, false, *envFile)

	gormDB, err := newDB(a.config, a.logger)
	if err != nil {
		a.logger.WithError(err).Fatal("Unable to establish database connection")
	}

	db, err := persistence.New(
		relational.NewRelationalDAL(gormDB),
	)
	if err != nil {
		a.logger.WithError(err).Fatal("Unable to create persistence layer")
	}

	if a.config.App.SingleNode {
		if err := db.Migrate(); err != nil {
			a.logger.WithError(err).Fatal("Error applying database migrations")
		} else {
			a.logger.Info("Successfully applied database migrations")
		}
	}

	fs := public.NewLocalizedFS(a.config.App.Locale.String())
	gettext, gettextErr := locales.GettextFor(a.config.App.Locale.String())
	if gettextErr != nil {
		a.logger.WithError(gettextErr).Fatal("Failed reading locale files, cannot continue")
	}
	tpl, tplErr := fs.HTMLTemplate(gettext)
	if tplErr != nil {
		a.logger.WithError(tplErr).Fatal("Failed parsing template files, cannot continue")
	}
	emails, emailErr := fs.EmailTemplate(gettext)
	if emailErr != nil {
		a.logger.WithError(emailErr).Fatal("Failed parsing template files, cannot continue")
	}

	mailer, err := a.config.NewMailer()
	if err != nil {
		a.logger.WithError(err).Fatal("Failed to initialize mailer")
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", a.config.Server.Port),
		Handler: router.New(
			router.WithDatabase(db),
			router.WithLogger(a.logger),
			router.WithTemplate(tpl),
			router.WithEmails(emails),
			router.WithConfig(a.config),
			router.WithFS(fs),
			router.WithMailer(mailer),
		),
	}
	go func() {
		if a.config.Server.SSLCertificate != "" && a.config.Server.SSLKey != "" {
			err := srv.ListenAndServeTLS(a.config.Server.SSLCertificate.String(), a.config.Server.SSLKey.String())
			if err != nil && err != http.ErrServerClosed {
				a.logger.WithError(err).Fatal("Error binding server to network")
			}
		} else if len(a.config.Server.AutoTLS) != 0 {
			m := autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				HostPolicy: autocert.HostWhitelist(a.config.Server.AutoTLS...),
				Cache:      autocert.DirCache(a.config.Server.CertificateCache),
				Email:      a.config.Server.LetsEncryptEmail,
			}
			go http.ListenAndServe(":http", m.HTTPHandler(nil))
			if err := http.Serve(m.Listener(), srv.Handler); err != nil && err != http.ErrServerClosed {
				a.logger.WithError(err).Fatal("Error binding server to network")
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				a.logger.WithError(err).Fatal("Error binding server to network")
			}
		}
	}()
	if len(a.config.Server.AutoTLS) != 0 {
		a.logger.Info("Server now listening on port 80 and 443 using AutoTLS")
	} else {
		a.logger.Infof("Server now listening on port %d", a.config.Server.Port)
	}

	if a.config.App.SingleNode {
		hourlyJob := time.Tick(time.Hour)
		runOnInit := make(chan bool)
		go func() {
			for {
				select {
				case <-hourlyJob:
				case <-runOnInit:
				}
				affected, err := db.Expire(config.EventRetention)
				if err != nil {
					a.logger.WithError(err).Errorf("Error pruning expired events")
					return
				}
				a.logger.WithField("removed", affected).Info("Cron successfully pruned expired events")
			}
		}()
		runOnInit <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		a.logger.WithError(err).Fatal("Error shutting down server")
	}

	a.logger.Info("Gracefully shut down server")
}