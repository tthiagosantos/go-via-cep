package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /viacep/{cep}", SearchCEP)
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SearchCEP(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reqCep, err := getCep(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(reqCep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getCep(cep string) (*ViaCep, error) {
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var viaCep *ViaCep
	err = json.Unmarshal(body, &viaCep)
	if err != nil {
		return nil, err
	}
	return viaCep, nil
}
