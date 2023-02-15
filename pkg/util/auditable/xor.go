package auditable

func Xor(as ...[32]byte) [32]byte {
	var m [32]byte
	for _, a := range as {
		for i := range a {
			m[i] = m[i] ^ a[i]
		}
	}
	return m
}

var nilId = [32]byte{}

func isNil(id [32]byte) bool {
	return id == nilId
}
