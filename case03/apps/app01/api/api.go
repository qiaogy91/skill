package api

import (
	"case03/apps/app01"
	"github.com/gin-gonic/gin"
	"github.com/qiaogy91/ioc"
	iocgin "github.com/qiaogy91/ioc/config/gin"
	"github.com/qiaogy91/ioc/config/log"
	"log/slog"
	"math/rand"
	"time"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
}

func (h *Handler) Name() string {
	return app01.AppName
}
func (h *Handler) Priority() int {
	return 499
}
func (h *Handler) Init() {
	h.log = log.Sub(app01.AppName)

	r := iocgin.ModuleRouter(h)
	r.GET("ping", h.Pong)
}

func (h *Handler) Pong(ctx *gin.Context) {
	rand.NewSource(time.Now().UnixNano())
	delay := 0.1 + rand.Float64()*(0.7-0.1)
	time.Sleep(time.Duration(delay * float64(time.Second)))

	ctx.JSON(200, gin.H{"msg": "pong"})
}

func init() {
	ioc.Api().Registry(&Handler{})
}
