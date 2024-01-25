package v2

import (
	"fmt"
	"strings"
)

type router struct {
	// trees 是按照 HTTP 方法来组织的
	// 如 GET => *node
	trees map[string]*node
}

type HandlerFunc func(ctx *Context)

type node struct {
	path string
	//children 子节点
	//子节点的 path => *node
	children map[string]*node
	//handler 命中路由之后的逻辑
	handler HandlerFunc
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}

// addRoute 注册路由
// method 是 HTTP 方法
//
//	path 必须以 / 开始并且结尾不能有 /，中间也不允许有连续的 /
func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	// 空字符串判断
	if path == "" {
		panic("web: 路由是空字符串")
	}
	// 必须以 / 为开头判断
	if path[0] != '/' {
		panic("web: 路由必须以 / 开头")
	}
	// 结尾不能有 /
	if path != "/" && path[len(path)-1] == '/' {
		panic("web: 不能以 / 为结尾")
	}

	root, ok := r.trees[method]
	//这是一个全新的 HTTP 方法，我们必须创建根节点
	if !ok {
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	// 路径冲突
	if path == "/" {
		if root.handler != nil {
			panic("web: 路径冲突")
		}
		root.handler = handler
		return
	}

	// 对路由进行切割
	segs := strings.Split(path[1:], "/")

	//开始进行处理
	for _, s := range segs {
		// 对空路径进行判断
		if s == "" {
			panic(fmt.Sprintf("web: 非法路由。不允许使用 //a/b, /a//b 之类的路由, [%s]", path))
		}
		root = root.childrenOfCreate(s)
	}
	if root.handler != nil {
		panic(fmt.Sprintf("web: 路由冲突[%s]", path))
	}
}

func (n *node) childrenOfCreate(path string) *node {
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[path]
	if !ok {
		child = &node{path: path}
		n.children[path] = child
	}
	return child
}
