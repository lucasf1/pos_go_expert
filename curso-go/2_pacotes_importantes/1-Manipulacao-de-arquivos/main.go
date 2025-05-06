package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Cria um arquivo
	f, err := os.Create("./arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escreve no arquivo
	// tamanho, err := f.WriteString("Hello, World!") 				// Escreve string
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo")) // Escreve bytes
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	// Abre o arquivo
	// arquivo, err := os.Open("./arquivo.txt")
	arquivo, err := os.ReadFile("./arquivo.txt") // Outra forma de abrir o arquivo
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("./arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Remover arquivo
	// err = os.Remove("./arquivo.txt")
	// if err != nil {
	// 	panic(err)
	// }
}
