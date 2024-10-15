package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func main() {
	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "svc01",                        // 通常设置为微服务名
			Subsystem: "user",                         // 模块名
			Name:      "user_total",                   // 指标名
			Help:      "the total user in current db", // 指标描述

			// 静态标签的key value
			ConstLabels: map[string]string{
				"k1": "v1",
				"k2": "v2",
			},
		},
		[]string{"k01", "k02"},
	)

	prometheus.MustRegister(vec)

	go func() {
		t := time.NewTicker(2 * time.Second)
		for range t.C {
			vec.WithLabelValues("value1", "value2").Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	app := gin.Default()
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
