package shamir

import (
	"math"
	"math/rand"
	"time"
)

var largePrime int64 = 82589933

type SecretSharer struct {
	secret       int64
	threshold    int64
	coefficients []int64
}

func New(secret int64, threshold int64) *SecretSharer {
	coefficients := make([]int64, threshold)
	rand.Seed(time.Now().UnixNano())
	for i := int64(0); i < threshold; i++ {
		coefficients[i] = rand.Int63n(largePrime)
	}

	return &SecretSharer{
		secret:       secret,
		threshold:    threshold,
		coefficients: coefficients,
	}
}

func (ss *SecretSharer) Compute(x int64) int64 {
	sum := ss.secret
	for i := int64(0); i < ss.threshold; i++ {
		sum += ss.coefficients[i] * int64(math.Pow(float64(x), float64(i+1)))
	}
	return sum
}
