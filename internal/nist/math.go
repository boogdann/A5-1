package nist

import "math"

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
