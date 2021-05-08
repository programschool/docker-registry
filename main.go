package main

import (
	"./config"
	"./router"
	"fmt"
	"github.com/labstack/echo"
	"os"
)

func main() {
	// Echo instance
	e := echo.New()
	initLog(e)
	// Routes
	router.Public(e)
	router.Private(e)
	conf := config.Load()
	// Start server
	fmt.Println("Private Registry Service For boxlayer.com")
	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	e.Logger.Fatal(e.StartTLS(address, conf.CertFile, conf.KeyFile))
}

func initLog(e *echo.Echo) {
	logFile, err := os.OpenFile("/home/logs/docker-registry.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// 设置存储位置
	e.Logger.SetOutput(logFile)
}
