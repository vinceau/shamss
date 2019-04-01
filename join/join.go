package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Tuple struct {
	x int64
	y int64
}

func main() {
	partsPtr := flag.String("parts", "", "The different parts")

	flag.Parse()

	if *partsPtr == "" {
		fmt.Println("No parts were given")
		return
	}

	s := strings.Split(*partsPtr, ";")
	var parts []Tuple
	parts = make([]Tuple, len(s))

	for i, part := range s {
		p := strings.Split(part, ",")
		x, err := strconv.Atoi(p[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		y, err := strconv.Atoi(p[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		parts[i] = Tuple{
			x: int64(x),
			y: int64(y),
		}
	}

	res := joinShares(parts)
	fmt.Printf("Secret: %v\n", res)
}

func joinShares(shares []Tuple) int64 {
	var sum int64
	for j, share := range shares {
		// fmt.Printf("%v, %v\n", j, share)
		var prod int64 = 1
		for m := 0; m < len(shares); m++ {
			if m != j {
				fmt.Printf("%v, %v\n", shares[m].x, shares[j].x)
				prod *= shares[m].x / (shares[m].x - shares[j].x)
			}
		}
		sum += share.y * prod
	}
	return sum
}
