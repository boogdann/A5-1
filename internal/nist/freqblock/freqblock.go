package freqblock

import "2/internal/nist"

type Test struct {
	bits      []byte
	blockSize int
}

func New(bits []byte, blockSize int) *Test {
	return &Test{
		bits:      bits,
		blockSize: blockSize,
	}
}

func (t *Test) Run() float64 {
	n := len(t.bits) / t.blockSize

	pis := make([]float64, n)
	for i := range pis {
		pis[i] = t.calcPi(i)
	}

	chi := t.calcChi(pis)
	pValue := nist.Igamc(float64(n)/2, chi/2)

	return pValue
}

func (t *Test) calcPi(prevNumber int) float64 {
	count := 0
	for i := 0; i < t.blockSize; i++ {
		count += int(t.bits[prevNumber*t.blockSize+i])
	}
	return float64(count) / float64(t.blockSize)
}

func (t *Test) calcChi(pis []float64) float64 {
	sum := 0.0
	for _, pi := range pis {
		sum = sum + (pi-0.5)*(pi-0.5)
	}

	return 4.0 * float64(t.blockSize) * sum
}
