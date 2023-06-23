package main

import "math"

//Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como parâmetro
//e calcule a área do círculo (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi.

type circulo struct {
	raio float64
}

func recebcirculos(c circulo) float64 {

	return math.Pi * c.raio * c.raio

}
