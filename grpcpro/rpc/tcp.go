package rpc

import (
	"encoding/binary"
	"net"
)

var numOfLengthBytes = 8

// ReadMsg reads a message from the connection.
func ReadMsg(conn net.Conn) ([]byte, error) {
	lenBs := make([]byte, numOfLengthBytes)
	_, err := conn.Read(lenBs)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint64(lenBs)
	data := make([]byte, length)
	_, err = conn.Read(data)
	return data, err
}

// EncodeMsg encodes the message.
func EncodeMsg(data []byte) []byte {
	reqLen := len(data)
	res := make([]byte, numOfLengthBytes+reqLen)
	binary.BigEndian.PutUint64(res[:numOfLengthBytes], uint64(reqLen))
	copy(res[numOfLengthBytes:], data)
	return res
}
