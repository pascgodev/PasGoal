package block

import "github.com/pasgo/pasgo/core/operation"

type Block struct {
	Header *Header
	OpList *[]operation.Operation
}
