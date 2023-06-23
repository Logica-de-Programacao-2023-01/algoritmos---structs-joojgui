package main

import (
	"fmt"
	"time"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    time.Time
	Preco   float64
}

func ViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {

		return Viagem{}
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {

	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC), Preco: 300.0},
		{Origem: "Rio de Janeiro", Destino: "Brasília", Data: time.Date(2023, time.June, 25, 0, 0, 0, 0, time.UTC), Preco: 450.0},
		{Origem: "São Paulo", Destino: "Salvador", Data: time.Date(2023, time.June, 30, 0, 0, 0, 0, time.UTC), Preco: 600.0},
	}

	viagemMaisCara := ViagemMaisCara(viagens)
	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.Origem)
	fmt.Println("Destino:", viagemMaisCara.Destino)
	fmt.Println("Data:", viagemMaisCara.Data)
	fmt.Println("Preço:", viagemMaisCara.Preco)
}
