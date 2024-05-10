package mods

import (
	"context"

	logging "data-collection-hub-server/pkg/zap"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"

	"data-collection-hub-server/pkg/prometheus"
)

type PrometheusMiddleware struct {
	prometheus *prometheus.Prometheus
	zap        *logging.Zap
}

func NewPrometheusMiddleware(prometheus *prometheus.Prometheus, zap *logging.Zap) *PrometheusMiddleware {
	return &PrometheusMiddleware{
		prometheus: prometheus,
		zap:        zap,
	}
}

func (p *PrometheusMiddleware) Register(app *fiber.App) error {
	if err := p.setupPath(app); err != nil {
		return err
	}
	app.Use(p.prometheus.PrometheusFiberHandler())
	return nil
}

func (p *PrometheusMiddleware) setupPath(app *fiber.App) error {
	ctx := context.Background()
	p.zap.SetTagInContext(ctx, logging.SystemTag)
	logger, err := p.zap.GetLogger(ctx)
	if err != nil {
		return err
	}
	logger.Info(
		"Prometheus middleware setup path", zap.String(
			"path",
			p.prometheus.PrometheusConfig.MetricPath,
		),
	)
	app.Get(
		p.prometheus.PrometheusConfig.MetricPath, func(c *fiber.Ctx) error {
			h := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
			h(c.Context())
			return nil
		},
	)
	return nil
}
