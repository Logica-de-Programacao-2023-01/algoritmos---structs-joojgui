package main

import "fmt"

type Funcionário struct {
	nome    string
	salário float64
	idade   int
}

func main() {
	funcionário := Funcionário{
		nome:    "João",
		salário: 3000.0,
		idade:   30,
	}

	fmt.Println("Salário atual:", funcionário.salário)

	aumentoPercentual := 10
	funcionário.aumentarSalário(aumentoPercentual)

	fmt.Println("Salário após aumento de", aumentoPercentual, "%:", funcionário.salário)

	reduçãoPercentual := 5
	funcionário.diminuirSalário(reduçãoPercentual)

	fmt.Println("Salário após redução de", reduçãoPercentual, "%:", funcionário.salário)

	tempoDeServiço := funcionário.calcularTempoDeServiço()
	fmt.Println("Tempo de serviço:", tempoDeServiço)
}

func (f *Funcionário) aumentarSalário(percentual int) {
	aumento := f.salário * float64(percentual) / 100
	f.salário += aumento
}

func (f *Funcionário) diminuirSalário(percentual int) {
	redução := f.salário * float64(percentual) / 100
	f.salário -= redução
}

func (f *Funcionário) calcularTempoDeServiço() int {
	idadeInicial := 18
	tempoDeServiço := f.idade - idadeInicial
	return tempoDeServiço
}
