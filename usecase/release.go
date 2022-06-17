package usecase

import (
	"math"
)

func primeNumber(n int) bool {
	if n < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Release(number int) bool {
	var result bool
	if primeNumber(number) {
		result = true
	} else {
		result = false
	}

	return result
}
