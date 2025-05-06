package main

const a = "Hello, World!"

var (
	b bool    = true
	c int     = 10
	d string  = "Lucas"
	e float64 = 1.2
)

func main() {

	// var a string

	b = true
	println(b)
	println(c)
	println(d)
	println(e)
	println(a)

	// b = 10 // erro - não consegue mudar o tipo

	// var a string = "X"
	a := "X"
	a = "fd"

	println(a)

	x()
}

func x() {
	println(a)
}
