package main

import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

// Pacotes e módulos parte 1, 2 e 3
// Instalando pacotes
// go get
// go mod tidy

func main() {
	s := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", s)

	fmt.Println(matematica.A)

	carro := matematica.Carro{
		Marca: "Fiat",
	}
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())
	fmt.Println(uuid.New())

}
