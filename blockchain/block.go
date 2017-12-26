package blockchain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type block struct {
	ID            string      `json:"id"`
	TimestampNano int64       `json:"timestamp_in_nanoseconds"`
	Data          interface{} `json:"data"`
	ProofOfWork   int64       `json:"proof_of_work"`
	PreviousHash  string      `json:"previous_hash"`
}

func newBlock(refB *block, pow int64, data interface{}) *block {
	newB := &block{
		ID:            uuid.NewV4().String(),
		TimestampNano: time.Now().UTC().UnixNano(),
		Data:          data,
		ProofOfWork:   pow,
	}

	if refB != nil {
		newB.PreviousHash = hashBlock(*refB)
	}

	return newB
}
