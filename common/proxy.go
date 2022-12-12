package common

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// SimpleProxy 简易的反向代理，将所有请求转发给其他端口
type SimpleProxy struct {
	Addr      string
	ProxyAddr string
}

// NewSimpleProxy 选定监听端口和代理端口创建新的代理
func NewSimpleProxy(port, proxyPort int) *SimpleProxy {
	return &SimpleProxy{
		Addr:      fmt.Sprint("127.0.0.1:", port),
		ProxyAddr: fmt.Sprint("http://127.0.0.1:", proxyPort),
	}
}

// Run 启动监听
func (r *SimpleProxy) Run() {
	log.Println("Starting http proxy at", r.Addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", r.SimpleHandler)

	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	go func() { log.Fatal(server.ListenAndServe()) }()
}

// SimpleHandler 代理服务
func (r *SimpleProxy) SimpleHandler(w http.ResponseWriter, req *http.Request) {
	// 1. 解析代理地址，并更改请求体的协议和主机
	proxy, _ := url.Parse(r.ProxyAddr)
	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host

	// 2. 请求下游
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(req)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	// 3. 返回上游
	for key, val := range resp.Header {
		for _, v := range val {
			w.Header().Add(key, v)
		}
	}
	bufio.NewReader(resp.Body).WriteTo(w)
}
