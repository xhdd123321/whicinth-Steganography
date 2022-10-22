// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/pprof"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	// your code ...
	steg := r.Group("/steg")
	{
		steg.POST("/encode_image", handler.EncodeImageFromImage)
		steg.POST("/decode_image", handler.DecodeImageFromImage)
		steg.POST("/encode_doc", handler.EncodeDocFromImage)
		steg.POST("/decode_doc", handler.DecodeDocFromImage)
		steg.POST("/decode_intelligent", handler.DecodeDocOrImageFromImage)
	}

	sys := r.Group("/sys")
	{
		sys.GET("/api_statistic", handler.GetApiStatistic)
		sys.GET("/monitor", handler.GetSysMonitor)
	}

	drift := r.Group("/drift")
	{
		drift.GET("/receive", handler.ReceiveDrift)
	}

	dashboard := r.Group("/admin")
	{
		pprof.RouteRegister(dashboard, "pprof")
	}
}
