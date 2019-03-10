package p2p

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	netData := make([]byte, 22)
	for {
		r := bufio.NewReader(c)
		_, err := r.Read(netData)
		if err != nil {
			fmt.Println(err)
			return
		}
		dataHeader := &DataHeader{}
		err = binary.Read(bytes.NewReader(netData[0:4]), binary.LittleEndian, &dataHeader.NetworkId)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = binary.Read(bytes.NewReader(netData[4:6]), binary.LittleEndian, &dataHeader.NetMethod)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = binary.Read(bytes.NewReader(netData[6:8]), binary.LittleEndian, &dataHeader.Operation)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = binary.Read(bytes.NewReader(netData[8:10]), binary.LittleEndian, &dataHeader.ErrCode)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = binary.Read(bytes.NewReader(netData[10:14]), binary.LittleEndian, &dataHeader.RequestId)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = binary.Read(bytes.NewReader(netData[14:16]), binary.LittleEndian, &dataHeader.ProtocolVersion)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = binary.Read(bytes.NewReader(netData[16:18]), binary.LittleEndian, &dataHeader.ProtocolAvailable)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = binary.Read(bytes.NewReader(netData[18:22]), binary.LittleEndian, &dataHeader.DataSize)
		if err != nil {
			fmt.Println(err)
			return
		}
		dataContent := make([]byte, dataHeader.DataSize)
		_, err = r.Read(dataContent)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Debug Begin
		fmt.Println(dataHeader)
		fmt.Println(dataContent)
		// Debug End
	}
}

// handleHello will handle the request whose requestMethod is NetRequest
func handleHello() {

}

func StartServer() {
	l, err := net.Listen("tcp4", ":5005")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
