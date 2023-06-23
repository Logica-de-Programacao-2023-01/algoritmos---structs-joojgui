package main

import "fmt"

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(index int) {
	if index >= 0 && index < len(a.Notas) {
		a.Notas = append(a.Notas[:index], a.Notas[index+1:]...)
	}
}

func (a *Aluno) CalcularMedia() float64 {
	if len(a.Notas) == 0 {
		return 0.0
	}

	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}

	media := soma / float64(len(a.Notas))
	return media
}

func (a *Aluno) ImprimirDados() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Média:", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{8.5, 9.0, 7.5},
	}

	aluno.ImprimirDados()

	aluno.AdicionarNota(6.0)
	fmt.Println("Notas após adicionar uma nova nota:", aluno.Notas)

	aluno.RemoverNota(1)
	fmt.Println("Notas após remover a segunda nota:", aluno.Notas)
}
