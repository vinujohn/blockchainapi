package blockchain

func mine(start int64) int64 {
	if start <= 0 {
		start = 1
	}

	var b []byte
	var i int64
	for i = 1; ; i++ {
		b = convertIntToBytes(start * i)

		if isValid(b) {
			return i
		}
	}
}
