package service

import (
	"context"
	"net/url"

	"github.com/khulnasoft/khulnasoft/lib/errors"

	"github.com/khulnasoft/khulnasoft/cmd/msp-example/internal/bigquery"
	"github.com/khulnasoft/khulnasoft/cmd/msp-example/internal/postgresql"
	"github.com/khulnasoft/khulnasoft/cmd/msp-example/internal/redis"
)

type serviceState struct {
	statelessMode bool

	bq *bigquery.Client
	rd *redis.Client
	pg *postgresql.Client
}

func (s serviceState) Healthy(ctx context.Context, _ url.Values) error {
	if s.statelessMode {
		return nil
	}

	// Write a single test event
	if err := s.bq.Write(ctx, "service.healthy"); err != nil {
		return errors.Wrap(err, "bigquery")
	}

	// Check redis connection
	if err := s.rd.Ping(ctx); err != nil {
		return errors.Wrap(err, "redis")
	}

	// Check database
	if err := s.pg.Ping(ctx); err != nil {
		return errors.Wrap(err, "postgresql")
	}

	return nil
}
