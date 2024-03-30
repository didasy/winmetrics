package app

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/didasy/winmetrics"
	"github.com/didasy/winmetrics/app/handler"
	"github.com/didasy/winmetrics/config"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

const (
	shutdownPeriod = 3 * time.Second
	runUpdater     = true
)

func Run(ctx context.Context, cfg *config.Configuration, log winmetrics.Logger) (err error) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT)
	defer stop()

	r := gin.New()

	r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	r.GET("/metrics", handler.Metrics(log, nil, runUpdater))

	apiGroup := r.Group("/api/v1")
	apiGroup.Use(gzip.Gzip(gzip.DefaultCompression))
	{
		apiGroup.GET("/metrics", handler.APIMetrics(log, nil, nil))
	}

	server := &http.Server{
		Addr:    cfg.App.ListenAddress,
		Handler: r,
	}
	errChan := make(chan error)
	go serve(server, errChan)

	select {
	case <-ctx.Done():
		stop()

		log.Infoln("Server shutting down, waiting 3 seconds to complete cleaning up")
		ctx, cancel := context.WithTimeout(context.TODO(), shutdownPeriod)
		defer cancel()

		err = server.Shutdown(ctx)
	case err = <-errChan:
	}

	return
}

func serve(server *http.Server, errChan chan<- error) {
	err := server.ListenAndServe()
	if err != nil {
		errChan <- err
	}
}
