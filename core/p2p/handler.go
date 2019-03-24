package p2p

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/glog"
	"net"
)

type Conn struct {
}

// handleConnection deal with all conn and read their msg
func (p *LocalPeer) handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	netData := make([]byte, 22)
	// getHeader
	for {
		r := bufio.NewReader(c)
		_, err := r.Read(netData)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Get NetData:", netData)
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

		switch dataHeader.NetMethod {
		case NetRequest:
			switch dataHeader.Operation {
			case OpHello:
				handleReqHello(c, dataContent, dataHeader.RequestId)
			case OpError:
				handleReqError(c, dataContent, dataHeader.RequestId)
			case OpMessage:
				handleReqMessage(c, dataContent, dataHeader.RequestId)
			case OpGetBlockHeaders:
				handleReqGetBlockHeaders(c, dataContent, dataHeader.RequestId)
			case OpGetBlocks:
				handleReqGetBlocks(c, dataContent, dataHeader.RequestId)
			case OpNewBlock:
				handleReqNewBlock(c, dataContent, dataHeader.RequestId)
			case OpNewBlock_FastPropagation:
				handleReqNewBlock_FastPropagation(c, dataContent, dataHeader.RequestId)
			case OpGetBlockChainOperations:
				handleReqGetBlockChainOperations(c, dataContent, dataHeader.RequestId)
			case OpAddOperations:
				handleReqAddOperations(c, dataContent, dataHeader.RequestId)
			case OpGetSafeBox:
				handleReqGetSafeBox(c, dataContent, dataHeader.RequestId)
			case OpGetPendingOperations:
				handleReqGetPendingOperations(c, dataContent, dataHeader.RequestId)
			case OpGetAccount:
				handleReqGetAccount(c, dataContent, dataHeader.RequestId)
			case OpGetPubKeyAccounts:
				handleReqGetPubKeyAccounts(c, dataContent, dataHeader.RequestId)
			case OpReservedStart:
				handleReqReservedStart(c, dataContent, dataHeader.RequestId)
			case OpReservedEnd:
				handleReqReservedEnd(c, dataContent, dataHeader.RequestId)
			case OpErrNotImpl:
				handleReqErrNotImpl(c, dataContent, dataHeader.RequestId)
			case NoOp:
			default:
				continue
			}
		case NetResponse:
			switch dataHeader.Operation {
			case OpHello:
				handleResHello(c, dataContent, dataHeader.RequestId)
			}
		}
	}
}

// handleHello will handle the request whose NetMethod is NetRequest and Operation is Hello
func (p *LocalPeer) handleHello(c net.Conn, dataContent []byte) {
	fmt.Println("Hello Received from", c.RemoteAddr())
}

// handleHello will handle the request whose NetMethod is NetRequest and Operation is Error
func (p *LocalPeer) handleError(c net.Conn, dataContent []byte) {
	fmt.Println("ERR: from: ", c.RemoteAddr(), " error: ", string(dataContent[:]))
}

func (p *LocalPeer) handleMessage(c net.Conn, dataContent []byte) {
	fmt.Println(c.RemoteAddr(), "says", string(dataContent[:]))
}

// handleReqHello will handleReq the request whose NetMethod is NetRequest and Operation is Hello
func handleReqHello(c net.Conn, dataContent []byte, requestId uint32) {
	// Hello ->
	/*
		<-
		{
			// Version info
			// Block & Bank & Treasury info
			// More
		}
	*/
	fmt.Println("Req Hello Received from", c.RemoteAddr())

	resBytes := GenBytes(NetResponse, OpHello, 0, 0, []byte{})
	l, err := c.Write(resBytes)
	if err != nil {
		glog.Warningln(err)
	}
	if l != len(resBytes) {
		handleReqHello(c, dataContent, requestId)
	}
}

func handleResHello(c net.Conn, dataContent []byte, requestId uint32) {
	fmt.Println("Res Hello Received from", c.RemoteAddr())
}

// handleReqHello will handleReq the request whose NetMethod is NetRequest and Operation is Error
func handleReqError(c net.Conn, dataContent []byte, requestId uint32) {
	fmt.Println("ERR: from: ", c.RemoteAddr(), " error: ", string(dataContent[:]))
}

func handleReqMessage(c net.Conn, dataContent []byte, requestId uint32) {
	fmt.Println(c.RemoteAddr(), "says", string(dataContent[:]))
}

// handleReqGetBlockHeaders
func handleReqGetBlockHeaders(c net.Conn, dataContent []byte, requestId uint32) {

}

func handleReqGetBlocks(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqNewBlock(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqNewBlock_FastPropagation(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqGetBlockChainOperations(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqAddOperations(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqGetSafeBox(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqGetPendingOperations(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqGetAccount(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqGetPubKeyAccounts(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqReservedStart(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqReservedEnd(c net.Conn, dataContent []byte, requestId uint32) {}

func handleReqErrNotImpl(c net.Conn, dataContent []byte, requestId uint32) {}
