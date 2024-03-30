package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/didasy/winmetrics"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func sensorsFuncTest(query string, dst interface{}, namespace string) (err error) {
	dv := dst.(*[]winmetrics.Sensor)
	*dv = append(*dv, winmetrics.Sensor{})

	return
}

func sensorsFuncTestFailed(query string, dst interface{}, namespace string) (err error) {
	err = errors.New("err")

	return
}

func hwFuncTest(query string, dst interface{}, namespace string) (err error) {
	dv := dst.(*[]winmetrics.Hardware)
	*dv = append(*dv, winmetrics.Hardware{})

	return
}

func hwFuncTestFailed(query string, dst interface{}, namespace string) (err error) {
	err = errors.New("err")

	return
}
func TestAPIMetrics(t *testing.T) {
	r := gin.Default()
	r.GET("/", APIMetrics(logrus.New(), hwFuncTest, sensorsFuncTest))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAPIMetricsFailedHW(t *testing.T) {
	r := gin.Default()
	r.GET("/", APIMetrics(logrus.New(), hwFuncTestFailed, sensorsFuncTestFailed))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAPIMetricsFailedSensors(t *testing.T) {
	r := gin.Default()
	r.GET("/", APIMetrics(logrus.New(), hwFuncTest, sensorsFuncTestFailed))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
