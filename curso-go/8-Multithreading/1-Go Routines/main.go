package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running.\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// thread 1
func main() {
	go task("A") // thread 2
	go task("B") // thread 3
	go func() {  // thread 4
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running.\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()
	// Nada aqui.
	// Sair
	time.Sleep(15 * time.Second)
}
