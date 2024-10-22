package main

import (
	_ "case04/apps"
	_ "github.com/qiaogy91/ioc/apps/health/restful"
	_ "github.com/qiaogy91/ioc/apps/metrics/restful" // metric
	_ "github.com/qiaogy91/ioc/apps/swagger/restful"
	_ "github.com/qiaogy91/ioc/config/cors/restful"
	_ "github.com/qiaogy91/ioc/config/otlp" // 开启遥测功能
	"github.com/qiaogy91/ioc/server"
)

func main() {
	server.Execute()
}
