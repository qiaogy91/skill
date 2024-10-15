package api

import (
	"github.com/emicklei/go-restful/v3"
	"log/slog"
	"math/rand"
	"time"
)

// Instrumentation libraries capture telemetry at the edges of your systems, such as inbound and outbound HTTP requests, but they don’t capture what’s going on in your application.
func (h *Handler) Pong(req *restful.Request, rsp *restful.Response) {
	ctx, span := h.tracer.Start(req.Request.Context(), "CustomSpanPong")
	defer span.End()

	rand.NewSource(time.Now().UnixNano())
	delay := 0.1 + rand.Float64()*(0.7-0.1)
	time.Sleep(time.Duration(delay * float64(time.Second)))

	u, err := h.svc.GetUser(req.Request.Context())
	if err != nil {
		if err := rsp.WriteAsJson(err); err != nil {
			h.log.Error("send json response err", slog.Any("err", err))
			return
		}
	}
	if err := rsp.WriteEntity(u); err != nil {
		h.log.Error("send json response err", slog.Any("err", err))
	}

	h.log.InfoContext(ctx, "received user query request")
}
