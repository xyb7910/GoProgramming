package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/silenceper/pool"
	"net"
	"reflect"
	"time"
)

func InitClientProxy(service Service, addr string) error {
	client, err := NewClient(addr)
	if err != nil {
		return err
	}
	return setFuncField(service, client)
}

// setFuncField  set service function field
func setFuncField(service Service, proxy Proxy) error {
	if service == nil {
		return errors.New("service is nil")
	}
	val := reflect.ValueOf(service)
	typ := val.Type()
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return errors.New("service is not a pointer to a struct")
	}

	val = val.Elem()
	typ = typ.Elem()

	numField := typ.NumField()
	for i := 0; i < numField; i++ {
		fieldTyp := typ.Field(i)
		fieldVal := val.Field(i)

		if fieldVal.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				// return value
				retVal := reflect.New(fieldTyp.Type.Out(0).Elem())
				// args [0]  is context.Context
				ctx := args[0].Interface().(context.Context)
				// args [1]  is *Request
				reqData, err := json.Marshal(args[1].Interface())
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}

				req := &Request{
					ServiceName: service.Name(),
					MethodName:  fieldTyp.Name,
					Arg:         reqData,
				}

				// call
				resp, err := proxy.Invoke(ctx, req)
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}

				err = json.Unmarshal(resp.Data, retVal.Interface())
				if err != nil {
					return []reflect.Value{retVal, reflect.ValueOf(err)}
				}
				return []reflect.Value{retVal, reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fnVal := reflect.MakeFunc(fieldTyp.Type, fn)
			fieldVal.Set(fnVal)
		}
	}
	return nil
}

type Client struct {
	pool pool.Pool
}

func NewClient(addr string) (*Client, error) {
	p, err := pool.NewChannelPool(&pool.Config{
		InitialCap:  1,
		MaxCap:      30,
		MaxIdle:     10,
		IdleTimeout: time.Second * 30,
		Factory: func() (interface{}, error) {
			return net.DialTimeout("tcp", addr, time.Second*3)
		},
		Close: func(conn interface{}) error {
			return conn.(net.Conn).Close()
		},
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		pool: p,
	}, nil
}

// Invoke service
func (c *Client) Invoke(ctx context.Context, req *Request) (*Response, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(data)
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: resp,
	}, nil
}

// Send data to server
func (c *Client) Send(data []byte) ([]byte, error) {
	val, err := c.pool.Get()
	if err != nil {
		return nil, err
	}
	conn := val.(net.Conn)
	defer func() {
		_ = conn.Close()
	}()
	req := EncodeMsg(data)
	_, err = conn.Write(req)
	if err != nil {
		return nil, err
	}
	return ReadMsg(conn)
}
