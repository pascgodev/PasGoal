package p2p

import (
	"bytes"
	"encoding/binary"
)

// genRequest will generate a binary ([]byte) message for request
func (p *Peer) genRequest(operation uint16, errorCode uint16, id uint32, dataContent []byte) []byte {
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
