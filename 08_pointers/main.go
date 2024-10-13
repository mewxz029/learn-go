package main

import "fmt"

func main() {
	// NOTE: same as var b *int
	b := new(int)
	var c **int

	a := 1
	// NOTE: assign reference of a (&) to b (pointer of int)
	b = &a
	// NOTE: assign reference of b (&) to c (pointer of pointer of int)
	c = &b

	fmt.Printf("address a %v\n", &a)
	fmt.Printf("address b %v\n", &b)
	fmt.Printf("address c %v\n", &c)

	fmt.Println()

	fmt.Printf("value a %v\n", a)
	fmt.Printf("value b %v\n", b)
	fmt.Printf("value c %v\n", c)

	fmt.Println()

	fmt.Printf("*b %v\n", *b)
	fmt.Printf("*c %v\n", *c)
	fmt.Printf("**c %v\n", **c)

	fmt.Println()

	n := 2
	nn := &n
	fmt.Println(n)
	// NOTE: pass by pointer
	double(nn)
	fmt.Println(n)

	// NOTE: pass by reference
	double(&n)
	fmt.Println(n)
}

func double(n *int) {
	*n *= 2
}
