package middleware

import (
	"github.com/gofiber/fiber"
)

func Proxy(app *fiber.App) {
	//app.Use(proxyTLS.New(proxyTLS.Config{
	//	Targets: []string{
	//		"192.168.20.101",
	//		//"127.0.0.1:3002",
	//	},
	//	//Rules: map[string]string{
	//	//	"/proxy": "/",
	//	//	"/3002": "/3002",
	//	//	"/3001": "/3001",
	//	//},
	//	Methods: []string{"GET"},
	//}))

	//app.Get("/v2", proxyTLS.Handler("192.168.20.101"))

	//app.Get("/v2/*", func(ctx *fiber.Ctx) {
	//	// Alter request
	//	ctx.Set("X-Forwarded-For", "3001")
	//	if err := proxyTLS.Forward(ctx, "192.168.20.101"); err != nil {
	//		ctx.SendStatus(503)
	//		ctx.Send(err.Error())
	//		return
	//	}
	//	// Alter response
	//	ctx.Set("X-Forwarded-By", "3001")
	//})

	//app.Get("/", func (c *fiber.Ctx) {
	//	c.Send("Hello, World!")
	//})

}
