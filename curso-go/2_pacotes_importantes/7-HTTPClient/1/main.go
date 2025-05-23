package main

import (
	"io"
	"net/http"
	"time"
)

func main() {

	// c := http.Client{Timeout: time.Microsecond} -> panic: Get "http://google.com": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
