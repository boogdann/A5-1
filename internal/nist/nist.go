package nist

import (
	"math"
)

func (t *NIST) Frequency() float64 {
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

func Igamc(a, x float64) float64 {
	const numSteps = 10000

	integrand := func(t float64) float64 {
		return math.Pow(t, a-1) * math.Exp(-t)
	}

	simpson := func(f func(float64) float64, a, b float64, n int) float64 {
		h := (b - a) / float64(n)
		sum := f(a) + f(b)
		for i := 1; i < n; i += 2 {
			sum += 4 * f(a+float64(i)*h)
		}
		for i := 2; i < n-1; i += 2 {
			sum += 2 * f(a+float64(i)*h)
		}
		return h / 3 * sum
	}

	result := simpson(integrand, 0, x, numSteps)

	return 1 - result/math.Gamma(a)
}

func (t *NIST) FrequencyBlock() float64 {
	n := len(t.bits) / t.blockCount

	pis := make([]float64, n)
	for i := range pis {
		pis[i] = t.calcPi(i)
	}

	chi := t.calcChi(pis)
	pValue := Igamc(float64(n)/2, chi/2)

	return pValue
}

func (t *NIST) calcPi(prevNumber int) float64 {
	count := 0
	for i := 0; i < t.blockCount; i++ {
		count += int(t.bits[prevNumber*t.blockCount+i])
	}
	return float64(count) / float64(t.blockCount)
}

func (t *NIST) calcChi(pis []float64) float64 {
	sum := 0.0
	for _, pi := range pis {
		sum = sum + (pi-0.5)*(pi-0.5)
	}

	return 4.0 * float64(t.blockCount) * sum
}

f