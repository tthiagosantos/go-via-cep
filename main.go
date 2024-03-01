package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

const url = "http://viacep.com.br/ws/69086581/json/"

func main() {

	fmt.Println(url)
	req, err := http.Get(url)
	if err != nil {
		log.Println("Error ao fazer requisicao: ", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(os.Stderr, "Erro ao ler resposta: ", err)
	}

	var data ViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		log.Println(os.Stderr, "Erro ao fazer parser da respotas", err)
	}
	fmt.Println(data)
}
