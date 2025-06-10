package main

import (
	"example.com/pos_go_expert/7_packaging/3/math"

	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(10, 20)
	println(m.Add())
	println(uuid.New().String())
}
