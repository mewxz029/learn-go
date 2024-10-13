package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}

	sumA := sum(a)
	sumB := sum(b)

	fmt.Println(sumA)
	fmt.Println(sumB)
}

type (
	Number interface {
		int | float64
	}

	Player[T Number] struct {
		Name   string
		Age    int
		Damage T
	}

	Database[T Number] interface{}
)

func sum[T Number](nums []T) T {
	var sum T

	for _, n := range nums {
		sum += n
	}

	return sum
}
