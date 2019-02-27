package beezap

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"go.uber.org/zap"
	"encoding/json"
)

func BeforeMiddlewareZap() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		ctx.Input.SetData("start_timer", time.Now())
	}
}

func AfterMiddlewareZap(logger *zap.Logger, timeFormat string, utc bool, appendBody bool) func(ctx *context.Context) {
	if appendBody {
		beego.Warn("[beezap] Be careful with personal data in body.")
	}

	return func(ctx *context.Context) {
		startTimeInterface := ctx.Input.GetData("start_timer")
		if startTime, ok := startTimeInterface.(time.Time); ok {
			path := ctx.Request.URL.Path
			query := ctx.Request.URL.RawQuery

			endTime := time.Now()
			latency := endTime.Sub(startTime)

			if utc {
				endTime = endTime.UTC()
			}

			headers, _ := json.Marshal(ctx.Request.Header)

			statusCode := ctx.Output.Status

			// TODO: The default code in beego is 0.
			if statusCode == 0 {
				statusCode = 200
			}

			fields := []zap.Field{
				zap.Int("status", statusCode),
				zap.String("method", ctx.Input.Method()),
				zap.String("path", path),
				zap.String("uri", ctx.Input.URI()),
				zap.String("query", query),
				zap.ByteString("headers", headers),
				zap.String("site", ctx.Input.Site()),
				zap.String("ip", ctx.Input.IP()),
				zap.String("refer", ctx.Input.Refer()),
				zap.String("user-agent", ctx.Input.UserAgent()),
				zap.String("time", endTime.Format(timeFormat)),
				zap.Duration("latency", latency),
			}

			if appendBody {
				fields = append(fields, zap.ByteString("body", ctx.Input.RequestBody))
			}

			logger.Info(path, fields...)
		}
	}
}

func InitBeeZapMiddleware(logger *zap.Logger, timeFormat string, utc bool,  appendBody ...bool) {
	beego.InsertFilter("*", beego.BeforeRouter, BeforeMiddlewareZap(), false)
	beego.InsertFilter("*", beego.FinishRouter, AfterMiddlewareZap(logger, timeFormat, utc, len(appendBody) > 0 && appendBody[0]), false)

	beego.Info("[beezap] Logger started")
}
