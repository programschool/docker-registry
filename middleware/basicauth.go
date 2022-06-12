// Package basicauth provides http basic authentication via middleware. See _examples/authentication/basicauth
package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/programschool/docker-registry/config"
	"github.com/programschool/docker-registry/library/request"
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
		user := map[string]interface{}{
			"username": username,
			"password": password,
		}
		client := request.Client{DeBug: conf.DeBug}.Create()

		data := new(verified)
		err := client.Post("registry/basic-auth", user).ParseJson(data)
		if err != nil {
			log.Println(err)
		}

		if data.Ok {
			return true, nil
		}
		return false, nil
	})
}
