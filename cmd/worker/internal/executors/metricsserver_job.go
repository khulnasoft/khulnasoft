package executors

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/khulnasoft/khulnasoft/cmd/worker/job"
	"github.com/khulnasoft/khulnasoft/internal/env"
	"github.com/khulnasoft/khulnasoft/internal/goroutine"
	"github.com/khulnasoft/khulnasoft/internal/httpserver"
	metricsstore "github.com/khulnasoft/khulnasoft/internal/metrics/store"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

type metricsServerJob struct{}

func NewMetricsServerJob() job.Job {
	return &metricsServerJob{}
}

func (j *metricsServerJob) Description() string {
	return "HTTP server exposing the metrics collected from executors to Prometheus"
}

func (j *metricsServerJob) Config() []env.Config {
	return []env.Config{metricsServerConfigInst}
}

func (j *metricsServerJob) Routines(_ context.Context, _ *observation.Context) ([]goroutine.BackgroundRoutine, error) {
	host := ""
	if env.InsecureDev {
		host = "127.0.0.1"
	}
	addr := net.JoinHostPort(host, strconv.Itoa(metricsServerConfigInst.MetricsServerPort))

	metricsStore := metricsstore.NewDistributedStore("executors:")

	handler := promhttp.HandlerFor(prometheus.GathererFunc(metricsStore.Gather), promhttp.HandlerOpts{})

	routines := []goroutine.BackgroundRoutine{
		httpserver.NewFromAddr(addr, &http.Server{Handler: handler}),
	}

	return routines, nil
}
