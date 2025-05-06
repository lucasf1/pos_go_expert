package main

// Quando usar Ponteiros

func soma(a, b int) int {
	a = 50
	return a + b
}

func soma2(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minhasVar1 := 10
	minhasVar2 := 20
	println(soma(minhasVar1, minhasVar2))
	println(minhasVar1)

	println(soma2(&minhasVar1, &minhasVar2))
	println(minhasVar1)

}
