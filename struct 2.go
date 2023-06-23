package main

import "fmt"

//Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço". O campo "endereço" deve ser uma outra struct
//com os campos "rua", "número", "cidade" e "estado". Escreva uma função que receba uma Pessoa como parâmetro
//e imprima seu endereço completo.

type endereco struct {
	Rua    string
	Numero int
	Estado string
	Cidade string
}

type pessoa struct {
	Nome     string
	idade    int
	Endereco endereco
}

func Enderecos(p pessoa) {
	fmt.Printf("Endereço completo de %s:\n", p.Nome)
	fmt.Printf("Rua: %s, Número: %d\n", p.Endereco.Rua, p.Endereco.Numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", p.Endereco.Cidade, p.Endereco.Estado)
}
