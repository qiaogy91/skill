package api

import (
	"case04/apps/user"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log    *slog.Logger
	svc    user.Service
	tracer trace.Tracer // tracer
}

func (h *Handler) Name() string  { return user.AppName }
func (h *Handler) Priority() int { return 499 }
func (h *Handler) Init() {
	h.log = log.Sub(user.AppName)
	h.svc = user.GetSvc()
	h.tracer = otel.GetTracerProvider().Tracer(user.AppName)

	// 路由定义
	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.GET("").To(h.Pong).
		Doc("获取用户").
		Metadata(restfulspec.KeyOpenAPITags, []string{"用户管理"}))
}

func init() {
	ioc.Api().Registry(&Handler{})
}
