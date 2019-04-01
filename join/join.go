package main

import (
	"flag"
	"fmt"

	"github.com/vinceau/shamss/shamir"
)

var largePrime int64 = 82589933

func main() {
	secretPtr := flag.Int("secret", 0, "The secret number to be split")
	sharesPtr := flag.Int("num_shares", 2, "The number of shares to generate")
	thresholdPtr := flag.Int("threshold", 1, "The minimum number of shares required to recreate the secret")

	flag.Parse()

	if *sharesPtr <= 0 {
		fmt.Printf("Shares must be greater than zero. Got: %v\n", *sharesPtr)
		return
	}

	if *thresholdPtr > *sharesPtr {
		fmt.Printf("Threshold must be greater than %v. Got: %v\n", *sharesPtr, *thresholdPtr)
		return
	}

	fmt.Println("shares: ", *sharesPtr)       // n
	fmt.Println("threshold: ", *thresholdPtr) // k

	generateShares(*secretPtr, *thresholdPtr, *sharesPtr)
}

func generateShares(secret int, threshold int, numShares int) error {
	ss := shamir.New(secret, threshold)

	for i := 0; i <= numShares; i++ {
		fmt.Printf("%v, %v\n", i+1, ss.Compute(i+1))
	}
	fmt.Println("\n")
	return nil
}
