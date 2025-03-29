package middleware

import (
	"api/configs/monitoring"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware(c *gin.Context) {
	path := c.Request.URL.Path
	method := c.Request.Method
	if path == "/metrics" {
		c.Next()
		return
	}
	defer func() {
		if err := recover(); err != nil {
			monitoring.RequestCountTotal.WithLabelValues().Inc()
			monitoring.RequestCount500.WithLabelValues(path, method, strconv.Itoa(http.StatusInternalServerError)).Inc()
			panic(err)
		}
	}()
	timer := prometheus.NewTimer(monitoring.RequestDuration.WithLabelValues(path, method))
	c.Next()
	status := c.Writer.Status()
	monitoring.RequestCountTotal.WithLabelValues().Inc()
	if status >= 400 {
		monitoring.RequestCount400.WithLabelValues(path, method, strconv.Itoa(status)).Inc()
		return
	}
	monitoring.RequestCount200.WithLabelValues(path, method, strconv.Itoa(status)).Inc()
	timer.ObserveDuration()
}
