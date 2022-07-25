package service

import "fmt"

//用于定义业务方法的接口
type IHelloService interface {
	Hello(name string) string
}

//用于实现上面定义的接口
type HelloService struct {
	//根据业务需求填充结构体
}

//实现上方定义的业务方法
func (h HelloService) Hello(name string) string {
	return fmt.Sprintf("Hello:%s", name)
}
