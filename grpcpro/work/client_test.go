package work

import (
	"net/rpc"
	"testing"
)

func TestDoClientWork(t *testing.T) {
	conn, _ := rpc.Dial("tcp", "127.0.0.1:1234")
	DoClientWork(conn)
}
