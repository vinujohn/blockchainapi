package blockchain

import (
	"fmt"
)

const (
	errInvalidHashFound = "block %s does not contain the correct hash for the previous block"
)

func Validate(bc *Blockchain) (bool, error) {
	var i, j int64

	c := bc.chain
	for i, j = 0, 1; j < bc.count; i, j = i+1, j+1 {
		if hashBlock(*c[i]) != c[j].PreviousHash {
			return false, fmt.Errorf(errInvalidHashFound, c[j].ID)
		}
	}

	return true, nil
}

func validateProofOfWork(bc *Blockchain, pow int64) bool {
	b := convertIntToBytes(bc.last.ProofOfWork * pow)

	return isValid(b)
}

// isValid will hash input data and determine its validity for the purposes of this blockchain
func isValid(b []byte) bool {
	h := hash(b)

	if h[0] == 0 && h[1] == 0 {
		return true
	}

	return false
}
