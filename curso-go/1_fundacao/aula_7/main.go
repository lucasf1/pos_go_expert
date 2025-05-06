package main

import "fmt"

func main() {
	salarios := map[string]int{"Lucas": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Lucas"])
	delete(salarios, "Lucas")
	fmt.Println(salarios["Lucas"])

	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"])

	sal := make(map[string]int)
	sal2 := map[string]int{}
	sal2["Lucas"] = 1000
	sal["Erika"] = 2000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)
	}

}
