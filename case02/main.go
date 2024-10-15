package main

import (
	"case02/exporter"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func main() {
	col := exporter.NewAppCollector()
	prometheus.MustRegister(col)

	go func() {
		tk := time.NewTicker(2 * time.Second)
		defer tk.Stop()

		for range tk.C {
			col.AddSample()
		}
	}()

	app := gin.Default()
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
