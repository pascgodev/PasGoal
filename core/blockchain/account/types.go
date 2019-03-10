package account

import (
	"crypto/ecdsa"
)

type AccountKey ecdsa.PublicKey

const (
	StateUnknown = iota
	StateNormal
	StateForSale
)

type Account struct {
	Account              uint32 // FIXED value. Account number
	AccountInfo          *AccountInfo
	Balance              uint64
	UpdatedBlock         uint32
	NOperation           uint32
	Name                 []byte
	AccountType          uint16
	PreviousUpdatedBlock uint32
}

type AccountInfo struct {
	State            int
	AccountKey       *AccountKey
	LockedUntilBlock uint32 // 0 = Not locked
	Price            uint64 // 0 = invalid price
	AccountToPay     uint32 // != itself
	NewPublicKey     *AccountKey
}

type OperationBlock struct {
	Block              uint32
	AccountKey         *AccountKey //TAccountKey
	Reward             uint64
	Fee                uint64
	ProtocolVersion    uint16
	ProtocolAvailable  uint16
	Timestamp          uint32
	CompactTarget      uint32
	Nonce              uint32
	BlockPayload       []byte
	InitialSafeboxHash []byte
	ProofOfWork        []byte
}

type BlockAccount struct {
	BlockChainInfo  *OperationBlock
	Accounts        [AccountsPerBlock]Account
	BlockHash       []byte
	AccumulatedWork uint64
}
