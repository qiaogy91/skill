package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
)

type AppCollector struct {
	QueueL prometheus.Gauge
	Total  prometheus.Counter
	Hist   prometheus.Histogram
	Sum    prometheus.Summary
}

func NewAppCollector() *AppCollector {
	return &AppCollector{
		QueueL: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "custom_queue_length",
		}),
		Total: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "custom_request_count",
		}),
		Hist: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "custom_request_duration",
			Buckets: []float64{10, 20, 30, 40, 50},
		}),
		Sum: prometheus.NewSummary(prometheus.SummaryOpts{
			Name: "custom_request_summary",
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.99: 0.001,
			},
		}),
	}
}

func (c *AppCollector) Describe(ch chan<- *prometheus.Desc) {
	c.QueueL.Describe(ch)
	c.Total.Describe(ch)
	c.Sum.Describe(ch)
	c.Hist.Describe(ch)
}

func (c *AppCollector) Collect(ch chan<- prometheus.Metric) {
	c.QueueL.Collect(ch)
	c.Total.Collect(ch)
	c.Sum.Collect(ch)
	c.Hist.Collect(ch)
}

func (c *AppCollector) AddSample() {
	// 模拟队列长度波动
	queueLength := float64(rand.Intn(100)) // 生成 0-99 的随机数
	c.QueueL.Set(queueLength)

	// 增加请求计数
	c.Total.Inc()

	// 模拟请求持续时间并记录直方图和摘要
	duration := float64(rand.Intn(100)) // 生成 0-99 的随机数
	c.Hist.Observe(duration)
	c.Sum.Observe(duration)
}
