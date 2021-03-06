package webserver

import (
	"strconv"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/tsingson/fasthttp-example/pkg/goutils"
)

func (ws *webServer) helloWorldGetHandler() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		tid := strconv.FormatInt(int64(ctx.ID()), 10)
		log := ws.Log.Named(tid)

		if ws.debug {
			log.Debug("helloWorldGetHandler")
			ctx.Request.Header.VisitAll(func(key, value []byte) {
				// log.Info("requestHeader", zap.String("key", gotils.B2S(key)), zap.String("value", gotils.B2S(value)))
				log.Debug(tid, zap.String("key", goutils.B2S(key)), zap.String("value", goutils.B2S(value)))
			})

			log.Debug(tid, zap.String("http payload", goutils.B2S(ctx.Request.Body())))

		}

		ctx.SetContentType(ContentText)
		ctx.SetStatusCode(200)
		ctx.SetBody([]byte(`hello world`))
		return
	}
}

func Hello() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		tid := strconv.FormatInt(int64(ctx.ID()), 10)

		ctx.Request.Header.Add("tid", tid)
		ctx.SetContentType(ContentText)
		ctx.SetStatusCode(200)
		ctx.SetBody([]byte(`hello world`))
		return
	}
}
