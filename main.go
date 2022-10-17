// Code generated by hertz generator.

package main

import (
	"math/rand"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/cronjob"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/godotenv"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/qiniu"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/redis"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	godotenv.InitGodotenv()
	viper.InitViper()
	redis.InitRedis()
	qiniu.InitQiniu()
	cronjob.InitCronjob()

	serverInit()
}

func serverInit() {
	h := server.Default(
		server.WithHostPorts(viper.Conf.App.HostPorts),
		server.WithMaxRequestBodySize(viper.Conf.App.MaxRequestBodySize),
	)
	register(h)
	h.Spin()
}
