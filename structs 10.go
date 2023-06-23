package main

import "fmt"

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverAvaliacao(index int) {
	if index >= 0 && index < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:index], f.Avaliacoes[index+1:]...)
	}
}

func (f *Filme) CalcularMediaAvaliacoes() float64 {
	if len(f.Avaliacoes) == 0 {
		return 0.0
	}

	soma := 0
	for _, avaliacao := range f.Avaliacoes {
		soma += avaliacao
	}

	media := float64(soma) / float64(len(f.Avaliacoes))
	return media
}

func (f *Filme) ImprimirInformacoes() {
	fmt.Println("Título:", f.Titulo)
	fmt.Println("Diretor:", f.Diretor)
	fmt.Println("Ano:", f.Ano)
	fmt.Println("Média de Avaliações:", f.CalcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		Titulo:     "Interstellar",
		Diretor:    "Christopher Nolan",
		Ano:        2014,
		Avaliacoes: []int{8, 9, 7, 10, 8},
	}

	filme.ImprimirInformacoes()

	filme.AdicionarAvaliacao(9)
	fmt.Println("Avaliações após adicionar uma nova avaliação:", filme.Avaliacoes)

	filme.RemoverAvaliacao(1)
	fmt.Println("Avaliações após remover a segunda avaliação:", filme.Avaliacoes)
}
