package rpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestInitClientProxy(t *testing.T) {
	server := NewServer()
	server.RegisterService(&UserService{})
	go func() {
		err := server.Start("tcp", ":8080")
		t.Log(err)
	}()
	time.Sleep(time.Second * 3)
	userClient := &UserService{}
	err := InitClientProxy(userClient, ":8080")
	require.NoError(t, err)
	resp, err := userClient.GetById(context.Background(), &GetByIdReq{Id: 123})
	assert.Equal(t, &GetByIdResp{
		Msg: "hello, world",
	}, resp)

}
