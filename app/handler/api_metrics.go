package handler

import (
	"net/http"

	"github.com/didasy/winmetrics"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func APIMetrics(log winmetrics.Logger, hwFunc, sensorsFunc winmetrics.QueryNamespaceFn) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		data := &winmetrics.Combined{}

		var err error
		data.Hardwares, err = winmetrics.GetHardwares(hwFunc)
		if err != nil {
			err = errors.Wrap(err, "failed to get hardwares from wmi")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		data.Sensors, err = winmetrics.GetSensors(sensorsFunc)
		if err != nil {
			err = errors.Wrap(err, "failed to get sensors from wmi")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, data)
	}

	return
}
