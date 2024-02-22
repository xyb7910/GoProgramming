package handler

type HelloServer struct{}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}
