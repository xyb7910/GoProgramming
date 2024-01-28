package v3

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type Context struct {
	Req        *http.Request
	Resp       http.ResponseWriter
	PathParams map[string]string

	//缓存的数据
	cacheQueryValue url.Values
}

type StringValue struct {
	val string
	err error
}

// 获取表单参数
func (c *Context) FormValue(key string) StringValue {
	if err := c.Req.ParseForm(); err != nil {
		return StringValue{
			err: err,
		}
	}
	return StringValue{val: c.Req.FormValue(key)}
}

// 获取请求参数
func (c *Context) QueryValue(key string) StringValue {
	if c.cacheQueryValue == nil {
		c.cacheQueryValue = c.Req.URL.Query()
	}

	vals, ok := c.cacheQueryValue[key]
	if !ok {
		return StringValue{err: errors.New("web: 找不到这个 key")}
	}
	return StringValue{val: vals[0]}
}

// 处理路径参数
func (c *Context) PathValue(key string) StringValue {
	val, ok := c.PathParams[key]
	if !ok {
		return StringValue{err: errors.New("web: 找不到这个 key")}
	}
	return StringValue{val: val}
}

// 设置cookie
func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Resp, cookie)
}

// 绑定 JSON
func (c *Context) BindJSON(val any) error {
	if c.Req.Body == nil {
		return errors.New("web: body 为 nil")
	}
	decoder := json.NewDecoder(c.Req.Body)
	return decoder.Decode(val)
}

// 状态码
func (c *Context) RespJSONOK(val any) error {
	return c.RespJSON(http.StatusOK, val)
}

func (c *Context) RespJSON(code int, val any) error {
	bs, err := json.Marshal(val)
	if err != nil {
		return err
	}
	c.Resp.WriteHeader(code)
	_, err = c.Resp.Write(bs)
	return err
}
