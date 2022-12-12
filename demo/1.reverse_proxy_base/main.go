package main

import (
	"gateway-demo/common"
	"os"
)

func main() {
	// 启动服务，监听 2003 端口并打印请求信息
	common.NewSimplseServer(2003).Run()

	// 启动代理，监听 2002 端口并转发给 2003 端口
	common.NewSimpleProxy(2002, 2003).Run()

	// 监听关闭信号
	forever := make(chan os.Signal)
	<-forever
}
