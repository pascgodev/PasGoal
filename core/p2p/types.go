package p2p

import (
	"github.com/pasgo/pasgo/core/blockchain/account"
)

type DataHeader struct {
	NetworkId         uint32
	NetMethod         uint16
	Operation         uint16
	ErrCode           uint16
	RequestId         uint32
	ProtocolVersion   uint16
	ProtocolAvailable uint16
	DataSize          uint32
}

type HelloDataContent struct {
	ServerPort         uint16
	PublicKey          string
	Timestamp          int32
	LastOpBlock        *account.OperationBlock
	NodeServerAddrList int32
}

type ErrorDataContent struct {
	ErrorData string
}

type MessageDataContent struct {
	MessageData string
}

type NodeServerAddress struct {
	Port              uint16
	LastConnTimestamp uint16
}
