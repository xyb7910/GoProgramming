package v2

import "net/http"

type Server interface {
	http.Handler

	//Start 启动服务器
	//addr 是监听地址， 如果需要指定端口，可以使用":8080" 或者 "localhost:8080"
	Start(addr string) error

	//addRoute 注册一个路由
	// method 是 HTTP 的方法
	addRoute(method string, path string, handler HandlerFunc)
	postRoute(method string, path string, handler HandlerFunc)
	// 我们并不采取这种设计方案
	// addRoute(method string, path string, handlers... HandleFunc)
}

type HTTPServer struct{ router }

var _Server = &HTTPServer{}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
}

// ServerHTTP HTTPServer 处理请求的入口
func (s *HTTPServer) Server(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	s.server(ctx)
}

// 启动服务器
func (s *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *HTTPServer) Get(path string, handler HandlerFunc) {
	s.addRoute(http.MethodGet, path, handler)
}

func (s *HTTPServer) Post(path string, handler HandlerFunc) {
	s.postRoute(http.MethodPost, path, handler)
}

func (s *HTTPServer) server(ctx *Context) {
	n, ok := s.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || n.handler == nil {
		ctx.Resp.WriteHeader(404)
		ctx.Resp.Write([]byte("Not Found"))
		return
	}
	n.handler(ctx)
}
