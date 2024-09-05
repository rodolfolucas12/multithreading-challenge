package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	model "github.com/rodolfolucas12/multithreading-challenge/cmd/model"
)

const (
	CEP_BUSCAR = "01310-200"
)

func main() {
	channelViaCep := make(chan model.ViaCep)
	channelBrasilApi := make(chan model.BrasilApi)

	go buscaViaCep(CEP_BUSCAR, channelViaCep)
	go buscaBrasilApi(CEP_BUSCAR, channelBrasilApi)

	select {
	case result := <-channelViaCep:
		fmt.Printf("VIACEP-API: %+v\n", result)
	case result := <-channelBrasilApi:
		fmt.Printf("BRASIL-API: %+v\n", result)
	case <-time.After(time.Second):
		fmt.Println("TIMEOUT - Tempo de resposta de 1 segundo atingido")
	}
}

func buscaViaCep(cep string, channel chan model.ViaCep) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	var endereco model.ViaCep

	err = json.Unmarshal(body, &endereco)
	if err != nil {
		fmt.Println(err)
		return
	}

	channel <- endereco
}

func buscaBrasilApi(cep string, channel chan model.BrasilApi) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	var endereco model.BrasilApi

	err = json.Unmarshal(body, &endereco)
	if err != nil {
		fmt.Println(err)
		return
	}

	channel <- endereco
}
