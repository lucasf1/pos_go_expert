package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processada com sucesso")     // Imprime no stdout
		w.Write([]byte("Request processada com sucesso")) // Escreve no Response
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
