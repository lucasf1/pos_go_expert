package main

import "fmt"

func main() {

	fmt.Println(sum(1, 12, 3, 45, 6, 34, 678, 534, 543))

	total := func() int {
		return sum(1, 12, 3, 45, 6, 34, 678, 534, 543) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
