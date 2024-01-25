package v1

import "net/http"

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler

	//Start 开启服务器
	//addr 是监听地址，如果只指定端口，可以使用":8080" 或者 ":localhost:8080"
	Start(addr string) error

	//addRoute 注册一个路由
	// method 是 HTTP 方法
	// path 为路径，必须以 / 开头
	AddRoute(method string, path string, handler HandleFunc)

	// 我们并不采取这种设计方案
	// addRoute(method string, path string, handlers... HandleFunc)
}

type HTTPServer struct{}

func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _Server = &HTTPServer{}

func (s *HTTPServer) ServerHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	s.server(ctx)
}
func (s *HTTPServer) server(ctx *Context) {
}

func (s *HTTPServer) AddRouter(method string, path string, handler HandleFunc) {
	panic("implement me")
}

func (s *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}
func (s *HTTPServer) Post(path string, handler HandleFunc) {
	s.AddRouter(http.MethodPost, path, handler)
}

func (s *HTTPServer) Get(path string, handler HandleFunc) {
	s.AddRouter(http.MethodPost, path, handler)
}
