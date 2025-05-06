package main

// Ponteiros

func main() {

	// Memória -> Endereço -> Valor
	a := 10
	println(a)
	println(&a)

	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	println(a)

	b := &a
	println(b)
	println(*b)
	*b = 30
	println(a)

}
