package rank

import (
	"2/internal/nist/matrix"
	"fmt"
	"math"
)

type Test struct {
	bits []byte
	q    int
	m    int
	n    int
}

func New(bits []byte, m, q int) *Test {
	return &Test{
		bits: bits,
		q:    q,
		m:    m,
		n:    len(bits),
	}
}

func (t *Test) Run() (float64, error) {
	minN := 38 * t.m * t.q
	if t.n < minN {
		return math.Inf(-1), fmt.Errorf("len(bits) < minN")
	}

	n := t.n / (t.m * t.q)
	m := min(t.m, t.q)

	matrices := t.createMatrices(n)
	ranks := t.createRanks(matrices, m)
	probs := t.probabilities(m)
	chi := t.chi(ranks, probs, n, m)

	pValue := math.Exp(-chi / 2)
	return pValue, nil
}

func (t *Test) createMatrices(n int) []matrix.BinMatrix {
	matrices := make([]matrix.BinMatrix, 0)
	for i := 0; i < n; i++ {
		matrices = append(matrices, matrix.New(t.bits, i, t.m, t.q))
	}

	return matrices
}

func (t *Test) createRanks(matrices []matrix.BinMatrix, m int) []float64 {
	ranks := make([]float64, m+1)

	for _, mx := range matrices {
		ranks[mx.Rank()]++
	}
	return ranks
}

func (t *Test) probabilities(m int) []float64 {
	probs := make([]float64, 0, m+1)
	for i := m; i >= 0; i-- {
		prob := math.Pow(2, float64(i*(t.q+t.m-i)-t.m*t.q)) * t.product(i)
		if prob < 0.0055 {
			break
		}
		probs = append(probs, prob)
	}

	probs[len(probs)-1] = 1 - sumTo(probs, len(probs)-1)
	return probs
}

func (t *Test) product(r int) float64 {
	p := 1.0
	for i := 0; i < r; i++ {
		p *= (1 - math.Pow(2, float64(i-t.q))) * (1 - math.Pow(2, float64(i-t.m))) / (1 - math.Pow(2, float64(i-r)))
	}
	return p
}

func (t *Test) chi(ranks []float64, probs []float64, n, m int) float64 {
	start := m - len(probs) + 2
	sum := 0.0
	for r := m; r >= start; r-- {
		sum += math.Pow(ranks[r]-float64(n)*probs[m-r], 2.0) / (float64(n) * probs[m-r])
	}
	sum += math.Pow(float64(n)-sumWith(ranks, start)-probs[len(probs)-1]*float64(n), 2.0) / (float64(n) * probs[len(probs)-1])
	return sum
}

func calculateSum(N float64, ranks []float64, startIndex int, probabilities []float64) float64 {
	sum := 0.0
	lastProb := probabilities[len(probabilities)-1]
	for _, rank := range ranks[startIndex:] {
		sum += math.Pow(N-rank-lastProb*N, 2) / (lastProb * N)
	}
	return sum
}

func sumWith(ranks []float64, startIndex int) float64 {
	sum := 0.0
	for _, rank := range ranks[startIndex:] {
		sum += rank
	}
	return sum
}

func sumTo(arr []float64, to int) float64 {
	sum := 0.0
	for _, v := range arr[:to] {
		sum += v
	}
	return sum
}
