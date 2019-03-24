package p2p

import (
	"bytes"
	"encoding/binary"
	"github.com/pasgo/pasgo/config"
	"net"
	"time"
	"github.com/golang/glog"
)

// genRequest will generate a binary ([]byte) message for request
func GenBytes(requestType uint16, operation uint16, errorCode uint16, id uint32, dataContent []byte) []byte {
	netId := make([]byte, 4)
	binary.LittleEndian.PutUint32(netId, NetId)

	reqType := make([]byte, 2)
	binary.LittleEndian.PutUint16(reqType, requestType)

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

// SendAllConnActiveHello send Hello to all conn
func (p *LocalPeer) SendAllConnActiveHello() {
	for _, conn := range p.ConnPool {
		p.SendHello(conn)
	}
}

// SendHello will send Hello to one conn
func (p *LocalPeer) SendHello(c *net.TCPConn) {
	/*
		HELLO command:
			- Operation stream
			- My Active server port (0 if no active). (2 bytes)
			- A Random Longint (4 bytes) to check if its myself connection to my server socket
			- My Unix Timestamp (4 bytes)
			- Registered node servers count
			(For each)
			- ip (string)
			- port (2 bytes)
			- last_connection UTS (4 bytes)
			- My Server port (2 bytes)
			- If this is a response:
			- If remote operation block is lower than me:
			- Send My Operation Stream in the same block thant requester
	*/

	port := make([]byte, 2)
	binary.LittleEndian.PutUint16(port, uint16(p.LocalPort))

	localKey := make([]byte, 4)
	binary.LittleEndian.PutUint32(localKey, uint32(p.LocalKey))

	timestamp := make([]byte, 4)
	binary.LittleEndian.PutUint32(timestamp, uint32(time.Now().Unix()))

	//TODO save block
	block := []byte{}

	//remotePeerNum := make([]byte, 4)
	//binary.LittleEndian.PutUint32(remotePeerNum, uint32(len(p.RemotePeers)))

	data := bytes.Join([][]byte{port, localKey, timestamp, block}, nil)

	remotePeerNum := make([]byte, 4)
	binary.LittleEndian.PutUint32(remotePeerNum, uint32(len(p.RemotePeers)))

	for _, peer := range p.RemotePeers {
		peerIP := []byte(peer.RemoteIP)

		peerPort := make([]byte, 2)
		binary.LittleEndian.PutUint16(peerPort, peer.RemotePort)

		lastTime := make([]byte, 4)
		binary.LittleEndian.PutUint32(lastTime, peer.LastTime)

		localPort := make([]byte, 2)
		binary.LittleEndian.PutUint16(localPort, peer.LocalPort)

		data = bytes.Join([][]byte{data, peerIP, peerPort, lastTime, localPort}, nil)
	}

	data = bytes.Join([][]byte{data, []byte(config.NodeConfig().NodeVersion)}, nil)
	// send accumulated work(???)

	b := GenBytes(NetRequest, OpHello, 0, 0, data)
	l, err := c.Write(b)
	if err != nil {
		glog.Warningln(err)
	}
	if l != len(b) {
		p.SendHello(c)
	}
}

func (p *LocalPeer) SendError(c *net.TCPConn, errString string) {
	data := []byte(errString)
	b := GenBytes(NetRequest, OpError, 0, 0, data)
	l, err := c.Write(b)
	if err != nil {
		glog.Warningln(err)
	}
	if l != len(b) {
		p.SendError(c, errString)
	}
}
