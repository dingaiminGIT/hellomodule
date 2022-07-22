package main

import (
	"github.com/dingaiminGIT/hellomodule/channel"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello,go module", zap.ByteString("url", ctx.RequestURI()))
}



func main() {
	//channel.TestChannel()
	//channel.Main()
	//channel.SelectMain()
	//channel.SelectMain2()
	//channel.LockMain()
	//fasthttp.ListenAndServe(":8081", fastHTTPHandler)
	channel.Lock2Main()
}
