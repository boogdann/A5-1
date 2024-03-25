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

func GammaUpperIncomplete(a, x float64) float64 {
	step := 0.001
	sum := 0.0

	for i := x; i < a; i += step {
		sum += math.Gamma(i) * step
	}

	return sum
}

func Igamc(a, x float64) float64 {
	gammaUpperIncomplete := GammaUpperIncomplete(a, x)
	gamma := math.Gamma(a)
	return gammaUpperIncomplete / gamma
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

func (t *NIST) CountSeries(bits []byte) float64 {
	seriesCount := 0
	currentBit := bits[0]
	seriesLength := 1

	for i := 1; i < len(bits); i++ {
		if bits[i] == currentBit {
			seriesLength++
		} else {
			if seriesLength > 1 {
				seriesCount++
			}
			currentBit = bits[i]
			seriesLength = 1
		}
	}

	if seriesLength > 1 {
		seriesCount++
	}

	return calculatePValue(seriesCount, len(bits))
}
func calculatePValue(seriesCount int, sequenceLength int) float64 {
	expectedMean := (sequenceLength + 1) / 2
	expectedVariance := (2*sequenceLength - 1) * (2*sequenceLength - 1) / (3 * sequenceLength)

	pValue := math.Erfc(math.Abs(float64(seriesCount)-float64(expectedMean)) / math.Sqrt(2*float64(expectedVariance)))

	return pValue
}
