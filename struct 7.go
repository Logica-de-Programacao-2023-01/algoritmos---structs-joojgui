package main

import "fmt"

type Animal struct {
	nome    string
	espécie string
	idade   int
	som     string
}

func main() {
	animal := Animal{
		nome:    "Rex",
		espécie: "Cachorro",
		idade:   5,
		som:     "Au Au",
	}

	fmt.Println("Informações do animal:")
	animal.imprimirInformações()

	fmt.Println("\nModificando o som do animal...")
	animal.modificarSom("Miau Miau")

	fmt.Println("\nNovas informações do animal:")
	animal.imprimirInformações()
}

func (a *Animal) modificarSom(novoSom string) {
	a.som = novoSom
}

func (a *Animal) imprimirInformações() {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Espécie: %s\n", a.espécie)
	fmt.Printf("Idade: %d anos\n", a.idade)
	fmt.Printf("Som: %s\n", a.som)
}
