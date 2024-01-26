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

// node 代表路由树的节点
// 路由树的匹配顺序为：
// 1、静态完全匹配
// 2、通配符匹配
// 这是不回溯匹配
type node struct {
	path string
	//children 子节点
	//子节点的 path => *node
	children map[string]*node
	//handler 命中路由之后的逻辑
	handler HandlerFunc

	//通配符 * 表示的节点，任意匹配
	starChild *node
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

func (r *router) findRoute(method string, path string) (*node, bool) {
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}

	if path == "/" {
		return root, true
	}

	segs := strings.Split(strings.Trim(path, "/"), "/")
	for _, s := range segs {
		root, ok = root.childof(s)
		if !ok {
			return nil, false
		}
	}
	return root, true
}

func (n *node) childof(path string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}
	res, ok := n.children[path]
	return res, ok
}
