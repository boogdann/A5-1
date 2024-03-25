package runs

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

	count := 0
	for i := 0; i < n; i++ {
		count += int(t.bits[i])
	}
	pi := float64(count) / float64(n)

	v := 1
	for i := 0; i < n-1; i++ {
		if !(t.bits[i] == t.bits[i+1]) {
			v++
		}
	}

	absVal := math.Abs(float64(v) - 2.0*pi*float64(n)*(1.0-pi))
	pValue := math.Erfc(absVal / (2 * math.Sqrt(2*float64(n)) * pi * (1.0 - pi)))

	return pValue
}
