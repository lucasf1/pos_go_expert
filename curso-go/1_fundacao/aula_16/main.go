package main

import "fmt"

// Ponteiros e Structs

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Lucas Farias"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{
		saldo: 0,
	}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {

	lucas := Cliente{
		nome: "Lucas",
	}
	lucas.andou()

	fmt.Println(lucas.nome)

	conta := Conta{
		saldo: 100,
	}
	conta.simular(200)
	println(conta.saldo)
}
