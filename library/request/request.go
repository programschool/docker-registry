package request

import (
	"../../config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
"Accept": "application/vnd.docker.distribution.manifest.v2+json",
*/

type Client struct {
	BaseUrl string
	ctx     context.Context
	Resp    *http.Response
	Headers map[string]string
	DeBug   bool
}

var headers = map[string]string{
	"Content-Type": "application/x-www-form-urlencoded",
	"Connection":   "Keep-Alive",
}

func (c Client) Create() Client {
	if len(c.BaseUrl) == 0 {
		c.BaseUrl = config.Load().Api
	}
	return c
}

// 发送普通数据请求 二维map
func (c Client) Post(url string, postData map[string]string) Client {
	targetUrl := c.getUrl(url)
	testHTTP(targetUrl)

	var r http.Request
	r.ParseForm()
	for k, v := range postData {
		r.Form.Add(k, v)
	}
	post := strings.TrimSpace(r.Form.Encode())
	req, err := http.NewRequest("POST", targetUrl, strings.NewReader(post))
	if err != nil {
		panic(err)
	}
	c.setHeader(req)
	resp, _ := http.DefaultClient.Do(req)
	c.Resp = resp

	c.deBug()
	return c
}

// 发送byte 数据请求
func (c Client) PostBytes(url string, postData interface{}) Client {
	targetUrl := c.getUrl(url)
	testHTTP(targetUrl)

	post, _ := json.Marshal(postData)
	req, err := http.NewRequest("POST", targetUrl, bytes.NewReader(post))
	if err != nil {
		panic(err)
	}
	c.setHeader(req)
	resp, _ := http.DefaultClient.Do(req)
	c.Resp = resp

	c.deBug()
	return c
}

func (c Client) Get(url string) Client {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	c.setHeader(req)
	resp, _ := http.DefaultClient.Do(req)
	c.Resp = resp

	c.deBug()
	return c
}

func (c Client) setHeader(req *http.Request) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
}

// 测试链接，防止拼写错误
func testHTTP(url string) {
	resp, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("\n\n\n%s: %s\n", url, resp.Status))
	}
}

// 获取绝对url地址
func (c Client) getUrl(url string) string {
	var targetUrl string
	if len(c.BaseUrl) != 0 {
		targetUrl = fmt.Sprintf("%s/%s", strings.Trim(c.BaseUrl, "/"), strings.Trim(url, "/"))
	} else {
		targetUrl = strings.Trim(url, "/")
	}

	return targetUrl
}

// 将数据解析为json
func (c Client) ParseJson(target interface{}) error {
	defer c.Resp.Body.Close()

	c.deBug()
	return json.NewDecoder(c.Resp.Body).Decode(&target)
}

func (c Client) deBug() {
	if c.DeBug == true {
		body, err := ioutil.ReadAll(c.Resp.Body)
		if err != nil {
			panic(err)
		}
		log.Println("\n\n\n\n======================DeBug======================")
		log.Println(string(body))
		log.Println("\n\n\n\n======================End DeBug======================")
		c.Resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
}
