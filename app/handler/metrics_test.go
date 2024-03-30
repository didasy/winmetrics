package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/didasy/winmetrics"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func metricsSensorsFuncTest(query string, dst interface{}, namespace string) (err error) {
	dv := dst.(*[]winmetrics.Sensor)
	*dv = append(*dv, winmetrics.Sensor{
		Name:       "test",
		SensorType: "test",
		Identifier: "test",
		Value:      1,
		Max:        2,
		Min:        0,
	})

	return
}

func TestMetrics(t *testing.T) {
	r := gin.Default()
	r.GET("/", Metrics(logrus.New(), metricsSensorsFuncTest, false))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMetricsFailed(t *testing.T) {
	r := gin.Default()
	r.GET("/", Metrics(logrus.New(), sensorsFuncTestFailed, false))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRunMetricsUpdate(t *testing.T) {
	prometheusSensors["test"] = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "test",
		Subsystem: "test",
		Name:      "test",
		Help:      "test",
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	runMetricsUpdate(ctx, logrus.New(), metricsSensorsFuncTest)

	// TODO: test get value of the gauge
}

func TestRunMetricsUpdateFailed(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	runMetricsUpdate(ctx, logrus.New(), sensorsFuncTestFailed)

	// TODO: test get value of the gauge
}
