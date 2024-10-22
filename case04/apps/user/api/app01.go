package api

import (
	"github.com/emicklei/go-restful/v3"
	"log/slog"
	"math/rand"
	"time"
)

/*
sum by (le) (rate(http_request_duration_histogram_bucket{path="/user/"}[5m]))
sum by (le) (rate(http_request_duration_histogram_bucket{path="/user/", method="GET"}[5m]))
sum by (le) (rate(http_request_duration_histogram_bucket{path="/user/", method="GET"}[5m]))
sum by (le) (rate(http_request_duration_histogram_bucket{path="/user/", service="case04"}[5m]))




*/

func (h *Handler) Pong(req *restful.Request, rsp *restful.Response) {
	ctx, span := h.tracer.Start(req.Request.Context(), "CustomSpanPong")
	defer span.End()

	rand.NewSource(time.Now().UnixNano())
	delay := 0.1 + rand.Float64()*(2)
	time.Sleep(time.Duration(delay * float64(time.Second)))

	u, err := h.svc.GetUser(ctx)
	if err != nil {
		if err := rsp.WriteAsJson(err); err != nil {
			h.log.Error("send json response err", slog.Any("err", err))
			return
		}
	}
	if err := rsp.WriteEntity(u); err != nil {
		h.log.Error("send json response err", slog.Any("err", err))
	}

	h.log.Info("这是没有context 的日志")
	h.log.InfoContext(ctx, "这是带有context 的日志")
}

// todo 不论使用的是 h.log.Info()  还是 h.log.InfoContext() ，日志数据都会发送给 Otlp 后端；唯一的区别是如果使用 h.log.InfoContext() 那么在 Trace 中能够关联到日志信息
