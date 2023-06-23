package main

import "fmt"

type Triângulo struct {
	base   float64
	altura float64
}

func main() {
	tri := Triângulo{base: 4.5, altura: 3.2}

	area := calcularÁreaTriângulo(tri)

	fmt.Printf("A área do triângulo é: %.2f\n", area)
}

func calcularÁreaTriângulo(t Triângulo) float64 {
	area := (t.base * t.altura) / 2
	return area
}
