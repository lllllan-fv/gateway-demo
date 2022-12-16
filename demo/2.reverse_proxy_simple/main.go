package main

import (
	"gateway-demo/common"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// 启动服务，监听 2003 端口并打印请求信息
	common.NewSimplseServer(2003).Run()

	reverse := "http://127.0.0.1:2003"
	url, _ := url.Parse(reverse)
	proxy := httputil.NewSingleHostReverseProxy(url)

	// 启动代理，监听 2002 端口并转发给 2003 端口
	log.Fatal(http.ListenAndServe("127.0.0.1:2002", proxy))
}
