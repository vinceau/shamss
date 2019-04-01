package shamir

import (
	"math"
	"math/rand"
)

type SecretSharer struct {
	secret       int
	threshold    int
	coefficients []float64
}

func New(secret int, threshold int) *SecretSharer {
	coefficients := make([]float64, threshold)
	for i := 0; i < threshold; i++ {
		coefficients[i] = rand.NormFloat64()
	}

	return &SecretSharer{
		secret:       secret,
		threshold:    threshold,
		coefficients: coefficients,
	}
}

func (ss *SecretSharer) Compute(x int) float64 {
	sum := float64(ss.secret)
	for i := 0; i < ss.threshold; i++ {
		sum += ss.coefficients[i] * math.Pow(float64(x), float64(i+1))
	}
	return sum
}
