package frequency

import "math"

type Test struct {
	bits []byte
}

func New(bits []byte) *Test {
	return &Test{
		bits: bits,
	}
}

func (t *Test) Run() float64 {
	n := len(t.bits)
	s := 0

	for _, bit := range t.bits {
		if bit == 1 {
			s++
		} else {
			s--
		}
	}

	sAbs := math.Abs(float64(s))
	result := math.Erfc(sAbs / math.Sqrt(2*float64(n)))

	return result
}
