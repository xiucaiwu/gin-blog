package main

import (
	"net/http"
	"fmt"
	//"log"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	//"github.com/fvbock/endless"
	//"syscall"
	//"os"
	//"os/signal"
	//"time"
	//"context"
	"gin-blog/pkg/logging"
	"gin-blog/models"
	"runtime"
)

func main()  {
	setting.Setup()
	models.Setup()
	logging.Setup()

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	if runtime.GOOS == "windows" {
		server := &http.Server{
			Addr:           endPoint,
			Handler:        routersInit,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		}

		server.ListenAndServe()
		return
	} else {
		//endless.DefaultReadTimeOut = readTimeout
		//endless.DefaultWriteTimeOut = writeTimeout
		//endless.DefaultMaxHeaderBytes = maxHeaderBytes
		//server := endless.NewServer(endPoint, routersInit)
		//server.BeforeBegin = func(add string) {
		//	log.Printf("Actual pid is %d", syscall.Getpid())
		//}
		//
		//err := server.ListenAndServe()
		//if err != nil {
		//	log.Printf("Server err: %v", err)
		//}
	}
}
