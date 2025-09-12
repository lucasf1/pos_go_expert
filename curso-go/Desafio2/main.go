package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type EnderecoBrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type EnderecoViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func WorkerBrasilApi(cep string, ch chan<- EnderecoBrasilApi) {
	// time.Sleep(time.Millisecond * 300)
	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}

	var data EnderecoBrasilApi
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
	}

	// fmt.Println(data)
	ch <- data
}

func WorkerViaBrasil(cep string, ch chan<- EnderecoViaCep) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}

	var data EnderecoViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
	}

	// fmt.Println(data)
	ch <- data
}

func main() {
	ch1 := make(chan EnderecoBrasilApi)
	ch2 := make(chan EnderecoViaCep)
	cep := "58042040"

	go WorkerViaBrasil(cep, ch2)
	go WorkerBrasilApi(cep, ch1)

	select {
	case end1 := <-ch1: // brasilApi
		fmt.Printf(
			"Recebido de brasilApi: Cep: %s, Estado: %s, Cidade: %s, Bairro: %s, Rua: %s\n",
			end1.Cep, end1.State, end1.City, end1.Neighborhood, end1.Street)

	case end2 := <-ch2: // viaCep
		fmt.Printf(
			"Recebido de viaCep: Cep: %s, Estado: %s, Cidade: %s, Bairro: %s, Rua: %s\n",
			end2.Cep, end2.Uf, end2.Localidade, end2.Bairro, end2.Logradouro)

	case <-time.After(1 * time.Second):
		println("timeout")
	}

}
