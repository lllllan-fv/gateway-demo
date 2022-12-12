package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// SimpleServer 简易的服务，请求任何地址都能够打印请求信息
type SimpleServer struct{ Addr string }

// NewSimplseServer 选定监听端口号创建新的服务
func NewSimplseServer(port int) *SimpleServer {
	return &SimpleServer{Addr: fmt.Sprint("127.0.0.1:", port)}
}

// Run 启动服务
func (r *SimpleServer) Run() {
	log.Println("Starting http server at", r.Addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", r.SimpleHandler)

	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	go func() { log.Fatal(server.ListenAndServe()) }()
}

// SimpleHandler 打印请求信息
func (r *SimpleServer) SimpleHandler(w http.ResponseWriter, req *http.Request) {
	upath := fmt.Sprint("http:/", "/", r.Addr, req.URL.Path, "\n")
	remoteAddr := fmt.Sprint("RemoteAddr = ", req.RemoteAddr, "\n")
	forward := fmt.Sprint("X-Forwarded-For = ", req.Header.Get("X-Forwarded-For"), "\n")
	realIp := fmt.Sprint("X-Real-Ip = ", req.Header.Get("X-Real-Ip"), "\n")
	header := fmt.Sprint("headers = ", req.Header, "\n")

	content := fmt.Sprint(upath, remoteAddr, forward, realIp, header)
	io.WriteString(w, content)
}
