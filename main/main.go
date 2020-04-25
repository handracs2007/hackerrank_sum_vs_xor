// https://www.hackerrank.com/challenges/sum-vs-xor/problem

package main

import (
	"fmt"
	"math"
)

func sumXor(n int64) int64 {
	var count int

	for n > 0 {
		if n%2 == 0 {
			count++
		}

		n /= 2
	}

	return int64(math.Pow(2.0, float64(count)))
}

func main() {
	fmt.Println(sumXor(4))
}
