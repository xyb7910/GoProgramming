package v2

import "testing"

func TestServer(t *testing.T) {
	s := NewHTTPServer()
	s.Get("/", func(ctx *Context) {
		ctx.Resp.Write([]byte("Hello Wrold"))
	})
	s.Get("/user", func(ctx *Context) {
		ctx.Resp.Write([]byte("Hello, user"))
	})
	s.Start(":8080")
}
