package auditable

func Xor(as ...Id) Id {
	var m Id
	for _, a := range as {
		for i := range a {
			m[i] = m[i] ^ a[i]
		}
	}
	return m
}

type Id [32]byte

var NilId Id = [32]byte{}

func IsNil(id Id) bool {
	return id == NilId
}
