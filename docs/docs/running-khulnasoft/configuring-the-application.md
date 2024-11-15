---
layout: default
title: Configuring the application
nav_order: 7
description: "How to configure an KhulnaSoft instance at runtime."
permalink: /running-khulnasoft/configuring-the-application/
parent: For operators
---

<!--
Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-License-Identifier: Apache-2.0
-->

# Configuring the application at runtime
{: .no_toc }

At runtime, __KhulnaSoft Fair Web Analytics is configured using environment variables__. All variables are following the pattern of `KHULNASOFT_<scope>_<key>` (e.g. `KHULNASOFT_SERVER_PORT`).

In addition to setting variables in the host environment __It also supports setting these values through [`env` files][dotenv]__.

[dotenv]: https://github.com/joho/godotenv

---

## Table of contents
{: .no_toc }

1. TOC
{:toc}

---

## Lookup order of .env files

### On Linux and MacOS
{: .no_toc }

In case the `-envfile` flag was supplied with a value when invoking a command, KhulnaSoft Fair Web Analytics will use this file. In case no such flag was given, it looks for files named `khulnasoft.env` in the following locations:

- In the current working directory
- In `~/.config`
- In `$XDG_CONFIG_HOME`
- In `/etc/khulnasoft`

### On Windows
{: .no_toc }

In case the `-envfile` flag was supplied with a value when invoking a command, KhulnaSoft Fair Web Analytics will use this file. In case no such flag was given, it expects a file named `khulnasoft.env` to be present in the current working directory.

## Configuration format

`env` files will specify the same keys the environment variables use, e.g.:

```
KHULNASOFT_SERVER_PORT="4000"
KHULNASOFT_DATABASE_DIALECT="sqlite3"
KHULNASOFT_DATABASE_CONNECTIONSTRING="/opt/khulnasoft/data/db.sqlite"
```

---

## Configuration options

__Heads Up__
{: .label .label-red }

All values for options can also be read from files (e.g. using Docker secrets).
To use this feature, use the option key suffixed with `_FILE` and set the path of the file containing the value,
e.g. `KHULNASOFT_DATABASE_CONNECTIONSTRING=/run/secrets/db_connection_string`.

### HTTP server

The `SERVER` namespace collects settings that affect the behavior of the HTTP server that is serving the application.

### KHULNASOFT_SERVER_PORT
{: .no_toc }

Defaults to `3000`.

The port the application listens on.

__Heads Up__
{: .label .label-red }

The Docker image sets this value to 80 in the Dockerfile, so you cannot override it from within an env file. Instead, map port 80 in the container to the desired port on your host system.

### KHULNASOFT_SERVER_REVERSEPROXY
{: .no_toc }

Defaults to `false`.

If set to `true` the application will assume it is running behind a reverse proxy. This means it does not add caching or security related headers to any response. Logging information about requests to `stdout` is also disabled.

### KHULNASOFT_SERVER_SSLCERTIFICATE
{: .no_toc }

In case you own a SSL certificate that is valid for the domain you are planning to serve your KhulnaSoft Fair Web Analytics instance from, you can pass the location of the certificate file using this variable. It also requires `KHULNASOFT_SERVER_SSLKEY` to be set.

### KHULNASOFT_SERVER_SSLKEY
{: .no_toc }

In case you own a SSL certificate that is valid for the domain you are planning to serve your KhulnaSoft Fair Web Analytics instance from, you can pass the location of the key file using this variable. It also requires `KHULNASOFT_SERVER_SSLCERTIFICATE` to be set.

### KHULNASOFT_SERVER_AUTOTLS
{: .no_toc }

In case you want KhulnaSoft Fair Web Analytics to automatically request a free SSL certificate from LetsEncrypt you can use this parameter and assign a comma separated list of supported domain names (e.g. `khulnasoft.mydomain.org,khulnasoft.otherdomain.org`) you are planning to serve it from. This will have the application automatically handle certificate issuing and renewal.

__Heads Up__
{: .label .label-red }

Using this feature will invalidate any port value that has been configured and will make KhulnaSoft Fair Web Analytics listen to both port 80 and 443. In such a setup, it is important that both ports are available to the public internet.

### KHULNASOFT_SERVER_CERTFICATECACHE
{: .no_toc }

Defaults to `/var/www/.cache` on Linux and MacOS, `%Temp%\khulnasoft.db` on Windows.

When using the AutoTLS feature, this sets the location where KhulnaSoft Fair Web Analytics will be caching certificates.

__Heads Up__
{: .label .label-red }

It is important that this value points to a persistent, non-ephemeral location as otherwise each request would issue a new certificate and your deployment will be rate limited by Let's Encrypt soon.

### KHULNASOFT_SERVER_LETSENCRYPTEMAIL
{: .no_toc }

In case you are using the AutoTLS feature, this setting can be used to pass an email to Let's Encrypt that will then be associated with the issued certificate. This allows Let's Encrypt to email you on certificate expiry or other possible issues with the certificate.

---

### Database

The `DATABASE` namespace collects settings regarding the connected persistence layer. If you do not configure any of these, KhulnaSoft Fair Web Analytics will be able to start, but data will not persist as it will be saved into a local temporary database.

### KHULNASOFT_DATABASE_DIALECT
{: .no_toc }

Defaults to `sqlite3`.

The SQL dialect to use. Supported options are `sqlite3`, `postgres` or `mysql`.

### KHULNASOFT_DATABASE_CONNECTIONSTRING
{: .no_toc }

Defaults to `/var/opt/khulnasoft/khulnasoft.db` on Linux and MacOS, `%Temp%\khulnasoft.db` on Windows.

The connection string or location of the database. For `sqlite3` this will be the location of the database file, for other dialects, it will be the URL the database is located at, __including the credentials__ needed to access it.

When using `mysql` make sure you append a `?parseTime=true` parameter to your connection string:

```
KHULNASOFT_DATABASE_CONNECTIONSTRING=user:pass@tcp(localhost:3306)/khulnasoft?parseTime=true
```

When using `postgres` and you are using a local database (or a Docker network) you might need to append a `?sslmode=disable` parameter to your connection string:

```
KHULNASOFT_DATABASE_CONNECTIONSTRING=postgres://user:pass@localhost:5432/khulnasoft?sslmode=disable
```

### KHULNASOFT_DATABASE_CONNECTIONRETRIES
{: .no_toc }

Defaults to `0`.

When running in a setup where you start the KhulnaSoft Fair Web Analytics server together with your database, you might run into race scenarios where KhulnaSoft Fair Web Analytics tries to connect to your database before it's ready to accept connections (e.g. docker-compose with MySQL). If needed, you can use this setting to tell it to retry connecting to the database after sleeping for a few seconds. This mechanism uses an exponential backoff algorithm, so if you specify a large number, the intervals might become big.

As this is more of a workaround, the __default behavior is not to retry__.

---

### Email

`SMTP` is a namespace used for configuring how transactional email is being sent. If any of these values is missing, KhulnaSoft Fair Web Analytics will fallback to using local `sendmail` which will likely be unreliable, so **configuring these values is highly recommended**.

### KHULNASOFT_SMTP_USER
{: .no_toc }

No default value.

The SMTP user name used when sending transactional email.

### KHULNASOFT_SMTP_PASSWORD
{: .no_toc }

No default value.

The SMTP user name used when sending transactional email.

### KHULNASOFT_SMTP_HOST
{: .no_toc }

No default value.

The SMTP hostname used when sending transactional email.

### KHULNASOFT_SMTP_PORT
{: .no_toc }

Default value `587`.

The SMTP port used when sending transactional email.

### KHULNASOFT_SMTP_SENDER
{: .no_toc }

Default value `no-reply@khulnasoft.com`.

The From address used when sending transactional email.

### KHULNASOFT_SMTP_AUTHTYPE
{: .no_toc }

Default value `LOGIN`.

The SMTP authentication type used when sending transactional email. Supported types are: `LOGIN`, `PLAIN`, `CRAM-MD5`, `NOAUTH`.

---

### Secrets

`KHULNASOFT_SECRET` is a single value.

### KHULNASOFT_SECRET
{: .no_toc }

No default value.

A Base64 encoded secret that is used for signing cookies and validating URL tokens. Ideally, it is of 16 bytes length. __If this is not set, a random value will be created at application startup__. This would mean that KhulnaSoft Fair Web Analytics can serve requests, but __an application restart would invalidate all existing sessions and all pending invitation/password reset emails__. If you do not want this behavior, populate this value, which is what we recommend.

---

__Heads Up__
{: .label .label-red }

The `khulnasoft` command has a `secret` subcommand you can use to generate such a value:

```
$ khulnasoft secret
INFO[0000] Created 16 bytes secret                       secret="NYOBGx2wF3CdrTva16m6BQ=="
```

Please __do not use the above example value__ when deploying your application.

---

### Application

The `APP` namespace affects how the application will behave.

### KHULNASOFT_APP_LOCALE
{: .no_toc }

Defaults to `en`.

The language the application will use when displaying user facing text. Right now, `en` (English), `de` (German), `fr` (French), `es` (Spanish), `pt` (Portuguese) and `vi` (Vietnamese) are supported. In case you want to contribute to KhulnaSoft Fair Web Analytics by adding a new language, [we'd love to hear from you][email].

[email]: mailto:admin@khulnasoft.com

### KHULNASOFT_APP_LOGLEVEL
{: .no_toc }

Defaults to `info`.

Specifies the application's log level. Possible values are `debug`, `info`, `warn`, `error`. If you use a level higher than `info`, access logging - which is happening at `info` level - will be suppressed.

### KHULNASOFT_APP_SINGLENODE
{: .no_toc }

Defaults to `true`.

In case you want to run KhulnaSoft Fair Web Analytics as a horizontally scaling service, you can set this value to `false`. This will disable all cron jobs and similar that handle automated database migration and event expiration.

### KHULNASOFT_APP_ROOTACCOUNT
{: .no_toc }

No default value.

If you want to collect usage statistics for your installation using KhulnaSoft Fair Web Analytics, you can use this parameter to specify an Account ID known to your KhulnaSoft Fair Web Analytics instance that will be used for collecting data.

### KHULNASOFT_APP_RETENTION
{: .no_toc }

Defaults to `6months`

By default, KhulnaSoft Fair Web Analytics retains data for 6 months (186 days) and deletes all data that is older than this threshold.
In case you wish to expire data even earlier, use this setting to define a shorter retention period.
Possible values are:

- `6months`
- `12weeks`
- `6weeks`
- `30days`
- `7days`

__Heads Up__
{: .label .label-red }

Please note that when you configure this value to be lower than what was usedbefore, __the application will delete all events older than the new value on startup__, and there will be __no way to recover this data__.
