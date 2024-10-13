package main

import "fmt"

func main() {
	// a := map[int]int{
	// 	1: 20,
	// 	2: 40,
	// 	3: 60,
	// }
	a := make(map[int]int)
	a[1] = 20

	if _, ok := a[1]; ok {
		fmt.Println("OK", a[1])
	} else {
		fmt.Println("NOT OK")
	}

	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
	}

	fmt.Println(graph)

	dfs(graph, 1, make(map[int]bool))
}

func dfs(graph map[int][]int, vertex int, visited map[int]bool) {
	visited[vertex] = true

	for _, v := range graph[vertex] {
		fmt.Printf("-> %d", v)
		dfs(graph, v, visited)
	}
}
