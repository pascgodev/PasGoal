package p2p

const (
	NetId = 0x0A043580 // 0x04000000 testnet

	NetRequest  = 0x0001
	NetResponse = 0x0002
	NetAutoSend = 0x0003

	NoOp                       = 0x0000
	OpHello                    = 0x0001
	OpError                    = 0x0002
	OpMessage                  = 0x0003
	OpGetBlockHeaders          = 0x0005
	OpGetBlocks                = 0x0010
	OpNewBlock                 = 0x0011
	OpNewBlock_FastPropagation = 0x0012
	OpGetBlockChainOperations  = 0x0013
	OpAddOperations            = 0x0020
	OpGetSafeBox               = 0x0021
	OpGetPendingOperations     = 0x0030
	OpGetAccount               = 0x0031
	OpGetPubKeyAccounts        = 0x0032
	OpReservedStart            = 0x1000
	OpReservedEnd              = 0x1FFF
	OpErrNotImpl               = 0x00FF

	NoError                     = 0x0000
	ErrorInvalidProtocolVersion = 0x0001
	ErrorIPBlackListed          = 0x0002
	ErrorNotFound               = 0x0003
	ErrorInvalidDataBufferInfo  = 0x0010
	ErrorInternalServerError    = 0x0011
	ErrorInvalidNewAccount      = 0x0012
	ErrorSafeBoxNotFound        = 0x0020
	ErrorNotAvailable           = 0x0021

	NoId = 0x0000

	ProtocolVersion   = 0x0009
	ProtocolAvailable = 0x0009
)
