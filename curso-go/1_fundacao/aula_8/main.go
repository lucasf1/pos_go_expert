package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(sum(1, 2))
	// fmt.Println(sum(51, 2))

	valor, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

// func sum(a int, b int) int {
// 	return a + b
// }

// func sum(a, b int) int {
// 	return a + b
// }

// func sum(a, b int) (int, bool) {
// 	if a+b > 50 {
// 		return a + b, true
// 	}
// 	return a + b, false
// }

func sum(a, b int) (int, error) {
	if a+b > 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
