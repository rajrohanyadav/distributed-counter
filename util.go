package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type intTuple struct {
	a, b int
}

func Zip(a, b []int) ([]intTuple, error) {

	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([]intTuple, len(a), len(b))

	for i, e := range a {
		r[i] = intTuple{e, b[i]}
	}

	return r, nil
}

func Sum(arr []int) int {
	sum := 0
	for _, e := range arr {
		sum += e
	}
	return sum
}
