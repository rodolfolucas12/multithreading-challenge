package model

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

type BrasilApi struct {
	Cep        string `json:"cep"`
	Estado     string `json:"state"`
	Cidade     string `json:"city"`
	Bairro     string `json:"neighborhood"`
	Logradouro string `json:"street"`
	Service    string `json:"service"`
}

type Response struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Uf         string `json:"uf"`
	Cidade     string `json:"cidade"`
}
