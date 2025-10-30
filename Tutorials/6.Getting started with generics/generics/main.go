package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

func main() {
	i := map[string]int64{"first": 1, "second": 2}

	f := map[string]float64{"first": 2., "second": 2.}
	fmt.Println("non-generics:", SumInts(i), SumFloats(f))
	fmt.Println("generics:", SumIntsOrFloats(i), SumIntsOrFloats(f))
}
