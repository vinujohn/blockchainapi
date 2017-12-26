package blockchain

import (
	"testing"
)

func TestHashAlgo(t *testing.T) {
	tests := []struct {
		desc         string
		expectedHash string
		testValue    block
	}{
		{
			"empty block",
			"2ae69f478ee1ff843f85cb181819974efb68289dadca867ad7871b1e1c60b150", //{"id":"","timestamp_in_nanoseconds":0,"data":null,"proof_of_work":0,"previous_hash":""}
			block{},
		},
		{
			"non-empty block",
			"c8206abde7d89f9349ab3bd71831fe871427b8888d132e83cc405232c451f1e0", //{"id":"23434","timestamp_in_nanoseconds":15000000000,"data":"Hello","proof_of_work":1000,"previous_hash":"630c06efbc7c1fecec41d82bd4d122d8761f829714dd8539910a95c780b10a00"}
			block{
				ID:            "23434",
				TimestampNano: 15000000000,
				Data:          "Hello",
				ProofOfWork:   1000,
				PreviousHash:  "630c06efbc7c1fecec41d82bd4d122d8761f829714dd8539910a95c780b10a00",
			},
		},
	}

	for _, test := range tests {
		if test.expectedHash != hashBlock(test.testValue) {
			t.Logf("test failed: %s", test.desc)
			t.FailNow()
		}
	}

}
