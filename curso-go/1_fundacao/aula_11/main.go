package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 44,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", lucas.Nome, lucas.Idade, lucas.Ativo)

	lucas.Ativo = false
	fmt.Println(lucas.Nome)
	fmt.Println(lucas.Ativo)

}
