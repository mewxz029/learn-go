package main

import "fmt"

func main() {
	// NOTE: Array (on stack mem)
	var a [3]int = [3]int{1, 2, 3}
	for _, v := range a {
		fmt.Println(v)
	}

	// NOTE: Slices (on heap mem)
	var b []int = []int{1, 2, 3}
	fmt.Println(b)

	c := make([]int, 0)
	c = append(c, 1, 2, 3)
	fmt.Println(c)
	c[2] = 5
	c = append(c, 4)
	fmt.Println(c)

	d := [3]int{1, 2, 3}
	fmt.Println("[Array] Inside main:")
	for i := range d {
		fmt.Printf("%v\n", &d[i])
	}
	double(d)
	fmt.Println(d) // NOTE: output: [1, 2, 3]

	f := []int{1, 2, 3}
	fmt.Println("[Slices] Inside main:")
	for i := range f {
		fmt.Printf("%v\n", &f[i])
	}
	doubleSlices(f)
	fmt.Println(f) // NOTE: output: [2, 4, 6]

	fmt.Println()

	g := []int{1, 2, 3}
	h := doubleSlicesReturn(g)
	fmt.Println(g)
	fmt.Println(h)
}

func double(nums [3]int) {
	fmt.Println("[Array] Inside double:")
	for i := range nums {
		fmt.Printf("%v\n", &nums[i])
		nums[i] *= 2
	}
}

func doubleSlices(nums []int) {
	// NOTE: slices that passed into function that using same address.
	fmt.Println("[Slices] Inside double:")
	for i := range nums {
		fmt.Printf("%v\n", &nums[i])
		nums[i] *= 2
	}
}

func doubleSlicesReturn(nums []int) []int {
	newNums := make([]int, len(nums))

	for i := range nums {
		newNums[i] = nums[i] * 2
	}

	return newNums
}
