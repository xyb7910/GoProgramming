package handler

const HelloServiceName = "handler/HelloServiceName"

type NewHelloService struct{}

func (s *NewHelloService) Hello(request string, reply *string) error {
	//返回值是通过修改replay的值
	*reply = "hello," + request
	return nil
}
