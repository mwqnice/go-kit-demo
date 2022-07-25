package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/mwqnice/go-kit-demo/service"
)
// Hello 业务使用的请求和响应格式
//HelloRequest 请求格式
type HelloRequest struct {
	Name string `json:"name"`
}

//HelloResponse响应格式
type HelloResponse struct {
	Reply string `json:"reply"`
}

//创建构造函数hello方法的业务处理
// MakeHelloServiceEndPorint 创建关于业务的构造函数
// 传入 Service/HelloService.go 定义的相关业务接口
// 返回 go-kit/endpoint.Endpoint (实际上就是一个函数签名)
func MakeHelloServiceEndPorint(s service.IHelloService) endpoint.Endpoint {
	// 这里使用闭包,可以在这里做一些中间件业务的处理
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		// request 是在对应请求来时传入的参数(这里的request 实际上是等下我们要将的Transport中一个decode函数中处理获得的参数)
		// 这里进行以下断言
		r, ok := request.(HelloRequest)
		if !ok {
			return HelloResponse{}, nil
		}
		// 这里实际上就是调用我们在Service/HelloService.go中定义的业务逻辑
		// 我们拿到了 Request.Name 那么我们就可以调用我们的业务 service.IHelloService 中的方法来处理这个数据并返回
		// 具体的业务逻辑具体定义....
		return HelloResponse{Reply: s.Hello(r.Name)}, nil
		// response 这里返回的response 可以返回任意的 不过根据规范是要返回我们刚才定义好的返回对象
	}
}