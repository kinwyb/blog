package controllers

import (
	"time"

	"github.com/rcrowley/go-metrics"
)

var enableMetrics = false //是否开启性能统计
var metricsRegistry metrics.Registry
var allRequestMetrics = "all.request"

//开启性能统计
func StartMetrics() {
	enableMetrics = true
	metricsRegistry = metrics.NewRegistry()
	//beego.Handler("/debug/metrics", exp.ExpHandler(metricsRegistry))
	//metrics.RegisterDebugGCStats(metricsRegistry)
	metrics.RegisterRuntimeMemStats(metricsRegistry)
	//go metrics.CaptureDebugGCStats(metricsRegistry, time.Second*5)
	go metrics.CaptureRuntimeMemStats(metricsRegistry, time.Second*5)
	//go influxdb.InfluxDB(metricsRegistry, time.Second, "http://118.31.188.107:18086", "rfid-test", "", "")
}

//结束性能统计
func StopMetrics() {
	enableMetrics = false
	metricsRegistry.UnregisterAll()
	metricsRegistry = nil
}
