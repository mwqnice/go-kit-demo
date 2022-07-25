package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/mwqnice/go-kit-demo/endpoint"
	"github.com/mwqnice/go-kit-demo/service"
	"github.com/mwqnice/go-kit-demo/transport"
	"net/http"
)

// 服务发布
func main()  {
	//1.创建我们最开始定义的service/HelloService.go
	s := service.HelloService{}

	//2.在用endpoint/HelloEndpoint创建业务服务
	hello := endpoint.MakeHelloServiceEndPorint(s)

	//3.使用kit创建handler
	//传入业务服务 以及 定义的加密解密方法
	helloServer := httpTransport.NewServer(hello, transport.HelloDecodeRequest,transport.HelloEncodeResponse)

	// 4.使用http包启动服务
	go http.ListenAndServe("0.0.0.0:8000", helloServer)
	select {}
}
