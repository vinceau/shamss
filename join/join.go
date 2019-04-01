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

	joinShares(parts)
}

func joinShares(shares []Tuple) error {
	for i, share := range shares {
		fmt.Printf("%v, %v\n", i, share)
	}
	return nil
}
