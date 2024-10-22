package main

import (
	_ "case03/apps"
	_ "github.com/qiaogy91/ioc/apps/health/gin"
	_ "github.com/qiaogy91/ioc/apps/metrics/gin"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/ioc/server"
)

func main() {
	server.Execute()
}
