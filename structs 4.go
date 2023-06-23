package main

import (
	"fmt"
	"time"
)

type Música struct {
	título  string
	artista string
	duração time.Duration
}

type Playlist struct {
	nome    string
	músicas []Música
}

func main() {
	playlist := Playlist{
		nome: "Minha Playlist",
		músicas: []Música{
			{título: "Música 1", artista: "Artista 1", duração: time.Minute * 3},
			{título: "Música 2", artista: "Artista 2", duração: time.Minute * 4},
			{título: "Música 3", artista: "Artista 3", duração: time.Minute * 5},
		},
	}

	imprimirPlaylist(playlist)
}

func imprimirPlaylist(p Playlist) {
	fmt.Println("Nome da Playlist:", p.nome)

	var tempoTotal time.Duration

	for _, música := range p.músicas {
		fmt.Println("Título:", música.título)
		fmt.Println("Artista:", música.artista)
		fmt.Println("Duração:", música.duração)

		tempoTotal += música.duração
	}

	fmt.Println("Tempo total da Playlist:", tempoTotal)
}
