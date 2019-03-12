package p2p

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type Conn struct {
}

// handleConnection deal with all conn and read their msg
func (p *Peer) handleConnection(c net.Conn) {
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
		fmt.Println(dataHeader.Operation)
		fmt.Println(dataHeader.DataSize)
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
				p.handleHello(c, dataContent)
			case OpError:
				p.handleError(c, dataContent)
			case OpMessage:
				p.handleMessage(c, dataContent)
			case OpGetBlockHeaders:
				p.handleGetBlockHeaders(c, dataContent)
			case OpGetBlocks:
				p.handleGetBlocks(c, dataContent)
			case OpNewBlock:
				p.handleNewBlock(c, dataContent)
			case OpNewBlock_FastPropagation:
				p.handleNewBlock_FastPropagation(c, dataContent)
			case OpGetBlockChainOperations:
				p.handleGetBlockChainOperations(c, dataContent)
			case OpAddOperations:
				p.handleAddOperations(c, dataContent)
			case OpGetSafeBox:
				p.handleGetSafeBox(c, dataContent)
			case OpGetPendingOperations:
				p.handleGetPendingOperations(c, dataContent)
			case OpGetAccount:
				p.handleGetAccount(c, dataContent)
			case OpGetPubKeyAccounts:
				p.handleGetPubKeyAccounts(c, dataContent)
			case OpReservedStart:
				p.handleReservedStart(c, dataContent)
			case OpReservedEnd:
				p.handleReservedEnd(c, dataContent)
			case OpErrNotImpl:
				p.handleErrNotImpl(c, dataContent)
			case NoOp:
			default:
				continue
			}
		}
	}
}

// handleHello will handle the request whose NetMethod is NetRequest and Operation is Hello
func (p *Peer) handleHello(c net.Conn, dataContent []byte) {
	fmt.Println("Hello Received from", c.RemoteAddr())
}

// handleHello will handle the request whose NetMethod is NetRequest and Operation is Error
func (p *Peer) handleError(c net.Conn, dataContent []byte) {
	fmt.Println("ERR: from: ", c.RemoteAddr(), " error: ", &dataContent)
}

func (p *Peer) handleMessage(c net.Conn, dataContent []byte) {

}

func (p *Peer) handleGetBlockHeaders(c net.Conn, dataContent []byte) {}

func (p *Peer) handleGetBlocks(c net.Conn, dataContent []byte) {}

func (p *Peer) handleNewBlock(c net.Conn, dataContent []byte) {}

func (p *Peer) handleNewBlock_FastPropagation(c net.Conn, dataContent []byte) {}

func (p *Peer) handleGetBlockChainOperations(c net.Conn, dataContent []byte) {}

func (p *Peer) handleAddOperations(c net.Conn, dataContent []byte) {}

func (p *Peer) handleGetSafeBox(c net.Conn, dataContent []byte) {}

func (p *Peer) handleGetPendingOperations(c net.Conn, dataContent []byte) {}

func (p *Peer) handleGetAccount(c net.Conn, dataContent []byte) {}

func (p *Peer) handleGetPubKeyAccounts(c net.Conn, dataContent []byte) {}

func (p *Peer) handleReservedStart(c net.Conn, dataContent []byte) {}

func (p *Peer) handleReservedEnd(c net.Conn, dataContent []byte) {}

func (p *Peer) handleErrNotImpl(c net.Conn, dataContent []byte) {}
