package nist

type NIST struct {
	bits       []byte
	blockCount int
}

func New(bits []byte, blockCount int) *NIST {
	return &NIST{
		bits:       bits,
		blockCount: blockCount,
	}
}
