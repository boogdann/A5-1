package discrete

import (
	"github.com/mjibson/go-dsp/fft"
	"math"
	"math/cmplx"
)

type Test struct {
	bits []byte
}

func New(bits []byte) *Test {
	return &Test{
		bits: bits,
	}
}

func (t *Test) Run() float64 {
	n := float64(len(t.bits))
	x := t.calcX()
	s := t.dft(x)
	m := t.calcM(s)
	tt := math.Sqrt(n * math.Log(1/0.05))
	n1 := float64(t.calcN1(m, tt))
	n0 := 0.95 * n / 2
	d := (n1 - n0) / math.Sqrt(n*0.95*0.05/4)

	pValue := math.Erfc(math.Abs(d) / math.Sqrt(2))
	return pValue
}

func (t *Test) calcX() []int {
	n := len(t.bits)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		if t.bits[i] == 1 {
			x[i] = 1
		} else {
			x[i] = -1
		}
	}

	return x
}

func (t *Test) dft(X []int) []complex128 {
	S := make([]complex128, len(X))
	for i, val := range X {
		S[i] = complex(float64(val), 0)
	}

	S = fft.FFT(S)

	return S
}

func (t *Test) calcM(S []complex128) []float64 {
	M := make([]float64, len(S)/2)
	for i := 0; i < len(S)/2; i++ {
		M[i] = cmplx.Abs(S[i])
	}
	return M
}

func (t *Test) calcN1(m []float64, tt float64) int {
	c := 0
	for _, val := range m {
		if val < tt {
			c++
		}
	}
	return c
}
