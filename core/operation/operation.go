package operation

import (
	"crypto/ecdsa"
	"github.com/pasgo/pasgo/core/blockchain/account"
)

const (
	Transaction = iota
	AutoBuyAccountTransaction
	BuyAccount
)

type OpTransactionData struct {
	Sender     uint32
	NOperation uint32
	Target     uint32
	Amount     uint64
	Fee        uint64
	Payload    []byte
	PublicKey  *ecdsa.PublicKey // .Sign (x,y) is in this obj

	// Protocol 2
	// Next values will only be filled after this operation is executed
	OperationType int
	AccountPrice  uint64
	SellerAccount uint32
	NewAccountKey *account.AccountKey
}

type OpChangeKeyData struct {
	AccountSigner uint32
	AccountTarget uint32
	NOperation    uint32
	Fee           uint64
	Payload       []byte
	PublicKey     *ecdsa.PublicKey // .Sign (x,y) is in this obj
	NewAccountKey account.AccountKey
}

type OpRecoverFoundsData struct {
	Account    uint32
	NOperation uint32
	Fee        uint64
}
