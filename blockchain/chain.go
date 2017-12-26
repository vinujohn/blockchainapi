package blockchain

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

const (
	errProofOfWorkInvalid = "proof of work is invalid"
)

type Blockchain struct {
	chain []*block
	last  *block
	count int64
}

func (bc *Blockchain) addBlock(b *block) {
	bc.chain = append(bc.chain, b)
	bc.last = b
	bc.count++
}

// New creates a new blockchain with an initial block
func New() *Blockchain {
	chain := &Blockchain{
		chain: []*block{},
	}

	// randomize initial proof of work
	rand.Seed(time.Now().UnixNano())
	pow := rand.Int63n(math.MaxInt64)

	firstB := newBlock(nil, pow, nil)

	chain.addBlock(firstB)

	return chain
}

func (bc *Blockchain) AddBlock(data interface{}, pow int64) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	newB := newBlock(bc.last, pow, data)

	if !validateProofOfWork(bc, pow) {
		return "", errors.New(errProofOfWorkInvalid)
	}

	bc.addBlock(newB)

	return bc.last.ID, nil
}
