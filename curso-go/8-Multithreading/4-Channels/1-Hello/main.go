package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // vazio

	// Thread 2
	go func() {
		canal <- "Olá mundo!" // cheio
	}()

	// Thread 1
	msg := <-canal // esvazia
	fmt.Println(msg)
}
