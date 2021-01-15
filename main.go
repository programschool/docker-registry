package main

import (
	"./config"
	"./router"
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()
	// Routes
	router.Public(e)
	router.Private(e)
	conf := config.Load()
	// Start server
	fmt.Println("Private Registry Service For boxlayer.com")
	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	e.Logger.Fatal(e.StartTLS(address, conf.CertFile, conf.KeyFile))
}
