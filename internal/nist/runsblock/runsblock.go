package runsblock

import (
	"2/internal/nist"
	"fmt"
	"math"
)

const (
	minN = 128
)

type Test struct {
	bits    []byte
	k       int
	m       int
	n       int
	pis     []float64
	initRun int
}

func New(bits []byte) *Test {
	var t Test
	n := len(bits)
	t.n = n
	t.pis = make([]float64, 0)
	t.bits = bits

	if n < 6272 {
		t.initRun = 1
		t.k = 3
		t.m = 8
		t.pis = append(t.pis, 0.21484375, 0.3671875, 0.23046875, 0.1875)
	} else if n < 75000 {
		t.initRun = 4
		t.k = 3
		t.m = 128
		t.pis = append(t.pis, 0.1174035788, 0.242955959, 0.249363483, 0.17517706, 0.102701071, 0.112398847)
	} else {
		t.initRun = 10
		t.k = 6
		t.m = 10000
		t.pis = append(t.pis, 0.0882, 0.2092, 0.2483, 0.1933, 0.1208, 0.0675, 0.0727)
	}

	return &t
}

func (t *Test) Run() (float64, error) {
	if len(t.bits) < minN {
		return math.Inf(-1), fmt.Errorf("len(bits) < minN")
	}

	vs := t.calcVs()
	chi := t.calcChi(vs)
	pValue := nist.Igamc(float64(t.n)/2, chi/2)

	return pValue, nil
}

func (t *Test) calcVs() []int {
	vs := make([]int, t.k+1)

	n := t.n / t.m
	for i := 0; i < n; i++ {
		longestRun := t.calcRun(i)
		if longestRun <= t.initRun {
			vs[0]++
		} else if longestRun >= t.initRun+t.k {
			vs[t.k]++
		} else {
			vs[longestRun-t.initRun]++
		}
	}

	return vs
}

func (t *Test) calcRun(prevBlocks int) int {
	run := 0
	longest := 0

	for i := 0; i < t.m; i++ {
		if t.bits[prevBlocks*t.m+i] == 1 {
			run++
			if run > longest {
				longest = run
			}
		} else {
			run = 0
		}
	}

	return longest
}

func (t *Test) calcChi(vs []int) float64 {
	chi := 0.0
	for i := 0; i < len(vs); i++ {
		chi += (float64(vs[i]) - float64(t.n)*t.pis[i]) *
			(float64(vs[i]) - float64(t.n)*t.pis[i]) / (float64(t.n) * t.pis[i])
	}
	return chi
}
