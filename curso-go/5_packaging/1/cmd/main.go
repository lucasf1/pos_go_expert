package main

import (
	"fmt"

	"example.com/pos_go_expert/7_packaging/1/math"
)

func main() {
	m := math.NewMath(10, 20)

	// m := math.Math{}

	fmt.Println("Math:", m)

	m.C = 30
	fmt.Println(m.C)


	result := m.Add()
	fmt.Println("Result:", result)
	fmt.Println(math.X)

}
