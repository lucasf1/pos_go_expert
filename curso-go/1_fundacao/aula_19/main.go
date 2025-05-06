package main

// Generics

/*
func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}
*/

type MyNumber int

// Constraint
type Number interface {
	~int | ~float64 // ~ para extender a relação (int ou tipos relacionados (MyNumber))
}

// func Soma[T int | float64](m map[string]T) T {
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false

}

func main() {
	m := map[string]int{"Lucas": 1000, "Erika": 2000, "Giovanna": 3000}
	// println(SomaInteiro(m))
	println(Soma(m))

	m2 := map[string]float64{"Lucas": 100.20, "Erika": 2000.3, "Giovanna": 300.0}
	// println(SomaFloat(m2))
	println(Soma(m2))

	m3 := map[string]MyNumber{"Lucas": 1000, "Erika": 2000, "Giovanna": 3000}
	println(Soma(m3))

	println(Compara(10, 10.0))
}
