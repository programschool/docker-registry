// Package basicauth provides http basic authentication via middleware. See _examples/authentication/basicauth
package middleware

import (
	"../config"
	"../library/request"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

type verified struct {
	Ok bool
}

var conf = config.Load()

/*
todo
[√] 通过接口验证
[-] 加入验证缓存【现在查询本地缓存，如果不存在或者错误，继续向接口查询】
*/
func Basicauth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		user := map[string]string{
			"username": username,
			"password": password,
		}
		client := request.Client{DeBug: conf.DeBug}.Create()

		data := new(verified)
		err := client.Post("registry/basic_auth", user).ParseJson(data)
		if err != nil {
			log.Println(err)
		}

		if data.Ok {
			return true, nil
		}
		return false, nil
	})
}
