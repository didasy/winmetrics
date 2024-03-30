package handler

import (
	"context"
	"strings"
	"time"

	"github.com/didasy/winmetrics"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Metrics(log winmetrics.Logger, sensorsFunc winmetrics.QueryNamespaceFn, runUpdater bool) (h gin.HandlerFunc) {
	setupMetrics(context.Background(), log, sensorsFunc, runUpdater)

	h = func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	}

	return
}

var (
	prometheusSensors = make(map[string]prometheus.Gauge, 10)
)

func setupMetrics(ctx context.Context, log winmetrics.Logger, sensorsFunc winmetrics.QueryNamespaceFn, runUpdater bool) {
	data := &winmetrics.Combined{}

	var err error

	data.Sensors, err = winmetrics.GetSensors(sensorsFunc)
	if err != nil {
		err = errors.Wrap(err, "failed to get sensors from wmi")
		log.Errorf("Prometheus sensors metrics error: %v", err)
		return
	}

	for _, s := range data.Sensors {
		id := strings.ReplaceAll(s.Identifier, "/", "_")
		id = strings.ReplaceAll(id, "-", "_")
		id = strings.ReplaceAll(id, "{", "_")
		id = strings.ReplaceAll(id, "}", "_")

		prometheusSensors[s.Identifier] = promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "Windows",
			Subsystem: "Sensors",
			Name:      id,
			Help:      "Referring to " + s.Name + " with sensor type " + s.SensorType,
		})
	}

	// Update every second
	if runUpdater {
		go runMetricsUpdate(ctx, log, sensorsFunc)
	}
}

func runMetricsUpdate(ctx context.Context, log winmetrics.Logger, sensorsFunc winmetrics.QueryNamespaceFn) {
	for {
		data := &winmetrics.Combined{}

		var err error

		data.Sensors, err = winmetrics.GetSensors(sensorsFunc)
		if err != nil {
			err = errors.Wrap(err, "failed to get sensors from wmi")
			log.Errorf("Prometheus sensors metrics error: %v", err)
			return
		}

		for _, s := range data.Sensors {
			prometheusSensors[s.Identifier].Set(float64(s.Value))
		}

		select {
		case <-time.After(time.Second):
		case <-ctx.Done():
			log.Infoln("Metrics updater exiting")
			return
		}

	}
}
