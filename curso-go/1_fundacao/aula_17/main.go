package main

import "fmt"

// Interfaces vazias

type x interface{}

func main() {

	var x interface{} = 10
	var y interface{} = "Hello world!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipode da variáve é %T e o valor é %v\n", t, t)
}
