package router

import (
	"../config"
	"../library/request"
	"../middleware"
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var conf = config.Load()

func Private(e *echo.Echo) {
	g := e.Group("/v2")

	g.Use(middleware.Basicauth())
	g.GET("", authedHandle)
	g.GET("/", authedHandle)
	g.HEAD("/:user/:name/manifests/:reference", privateProxyHandle)
	g.HEAD("/:user/:name/blobs/:digest", privateProxyHandle)
	g.PUT("/:user/:name/manifests/:reference", privateProxyHandle)
	g.DELETE("/:user/:name/manifests/:reference", privateProxyHandle)
	g.DELETE("/:user/:name/blobs/:digest", privateProxyHandle)
	g.POST("/:user/:name/blobs/uploads/", privateProxyHandle)
	g.PATCH("/:user/:name/blobs/uploads/:uuid", privateProxyHandle)
	g.PUT("/:user/:name/blobs/uploads/:uuid", privateProxyHandle)
	g.DELETE("/:user/:name/blobs/uploads/:uuid", privateProxyHandle)
	g.GET("/_catalog", authedHandle)
}

func Public(e *echo.Echo) {
	e.GET("/", homeHandle)
	g := e.Group("/v2")
	g.GET("/:user/:name/tags/list", publicProxyHandle)
	g.GET("/:user/:name/manifests/:reference", publicProxyHandle)
	g.GET("/:user/:name/blobs/:digest", publicProxyHandle)
	g.GET("/:user/:name/blobs/uploads/:uuid", publicProxyHandle)
}

func homeHandle(c echo.Context) error {
	//return c.String(http.StatusOK, "PS Hub")
	return c.Redirect(http.StatusMovedPermanently, "https://www.boxlayer.com/")
}

func publicProxyHandle(c echo.Context) error {
	// do someing
	return goProxy(c)
}

// 不需要认证即可访问
func authedHandle(c echo.Context) error {
	return goProxy(c)
}

// 认证并且拥有所有权利才可以访问
func privateProxyHandle(c echo.Context) error {
	username, _, _ := c.Request().BasicAuth()
	user := c.Param("user")
	name := c.Param("name")
	imageFullName := fmt.Sprintf("%s/%s", user, name)

	permit := checkPermit(c.Request().Host, username, imageFullName)

	if permit.Ok {
		manifest, digest := getManifest(c.Request())
		if manifest.SchemaVersion == 2 {
			var size int64
			for _, v := range manifest.Layers {
				size += v.Size
			}

			pathMap := strings.Split(c.Request().RequestURI, "/")
			tag := pathMap[len(pathMap)-1]

			data := map[string]interface{}{
				"username":      username,
				"imageFullName": imageFullName,
				"contentDigest": digest,
				"tag":           tag,
				"size":          size,
				"manifest":      manifest,
			}
			//log.Println(data)

			var requestTarget string
			for _, proxy := range conf.Proxy {
				if proxy.Match == c.Request().Host {
					requestTarget = proxy.Request
					break
				}
			}

			client := request.Client{DeBug: conf.DeBug}.Create()
			res := client.Post(requestTarget, data)
			body, _ := ioutil.ReadAll(res.Resp.Body)
			log.Println(string(body))
		}

		return goProxy(c)
	} else {
		return c.String(403, fmt.Sprintf(`{"msg": "%s"}`, permit.Message))
	}
}

type permit struct {
	Ok      bool
	Message string
}

func checkPermit(host string, username string, imageFullName string) *permit {
	client := request.Client{DeBug: conf.DeBug}.Create()
	post := map[string]interface{}{
		"host":          host,
		"username":      username,
		"imageFullName": imageFullName,
	}

	checked := new(permit)
	err := client.Post("registry/permit", post).ParseJson(checked)
	if err != nil {
		log.Println(err)
	}

	return checked
}

func getManifest(r *http.Request) (schema2.Manifest, string) {
	var manifest schema2.Manifest
	var digest string
	// application/vnd.docker.distribution.manifest.v1+prettyjws
	//  && r.Header.Get("Content-type") == "application/vnd.docker.distribution.manifest.v2+json"
	if r.ContentLength > 0 && r.Header.Get("Content-type") == "application/vnd.docker.distribution.manifest.v2+json" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		sum := sha256.Sum256(body)
		digest = fmt.Sprintf("sha256:%x", sum)

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, &manifest)
		if err != nil {
			log.Println(err)
		}
	}
	return manifest, digest
}

// Handler
func goProxy(c echo.Context) error {
	var target *url.URL
	for _, proxy := range conf.Proxy {
		if proxy.Match == c.Request().Host {
			target = &url.URL{
				Scheme: proxy.Scheme,
				Host:   fmt.Sprintf("%s:%s", proxy.Host, proxy.Port),
			}
			break
		}
	}

	if target == nil {
		return c.String(http.StatusOK, "没有匹配的服务")
	} else {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: target.Scheme,
			Host:   target.Host,
		})
		proxy.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		proxy.ServeHTTP(c.Response(), c.Request())
	}
	return nil
}
