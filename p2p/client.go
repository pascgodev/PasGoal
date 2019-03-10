package p2p

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
)

// Multi Client for each nodes

type Client struct {
	conn net.Conn
}

func NewP2PClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:4004")
	if err != nil {
		fmt.Println(err)
	}

	for {
		data, err := ioutil.ReadAll(conn)
		if len(data) > 0 {
			fmt.Println(string(data[:]))
		}
		if err != nil {
			fmt.Println("ERR")
			fmt.Println(err)
			//conn.Write(Req)
		}
	}
}

// genRequest will generate a binary ([]byte) message for request
func (c *Client) genRequest(operation uint16, errorCode uint16, id uint32, dataContent []byte) []byte {
	netId := make([]byte, 4)
	binary.LittleEndian.PutUint32(netId, NetId)

	reqType := make([]byte, 2)
	binary.LittleEndian.PutUint16(reqType, NetRequest)

	opType := make([]byte, 2)
	binary.LittleEndian.PutUint16(opType, operation)

	errCode := make([]byte, 2)
	binary.LittleEndian.PutUint16(errCode, errorCode)

	requestId := make([]byte, 4)
	binary.LittleEndian.PutUint32(requestId, id)

	protocolVersion := make([]byte, 2)
	binary.LittleEndian.PutUint16(protocolVersion, ProtocolVersion)

	protocolAvailable := make([]byte, 2)
	binary.LittleEndian.PutUint16(protocolAvailable, ProtocolAvailable)

	dataSize := make([]byte, 4)
	size := len(dataContent)
	binary.LittleEndian.PutUint32(dataSize, uint32(size))

	return bytes.Join([][]byte{
		netId,
		reqType,
		opType,
		errCode,
		requestId,
		protocolVersion,
		protocolAvailable,
		dataSize,
		dataContent,
	}, []byte{})
}
