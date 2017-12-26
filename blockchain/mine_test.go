package blockchain

import (
	"math"
	"testing"
)

func TestMiningAlgo(t *testing.T) {
	tests := []struct {
		desc          string
		startingProof int64
		expectedProof int64
	}{
		{
			"zero start proof",
			0,
			30556,
		},
		{
			"one start proof",
			1,
			30556,
		},
		{
			"two start proof",
			2,
			15278,
		},
		{
			"max start proof",
			math.MaxInt64,
			4663,
		},
	}

	for _, test := range tests {
		pow := mine(test.startingProof)
		if pow != test.expectedProof {
			t.Logf("actual proof of work %v not equal to expected proof of work %v", pow, test.expectedProof)
			t.FailNow()
		}
	}

}
