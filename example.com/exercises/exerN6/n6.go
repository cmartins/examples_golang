package main

import (
	"fmt"
)

func exe01() (int, string) {
	return 10, "Dez"
}

func exe02(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p1 pessoa) exe04() {
	fmt.Println("O nome da pessoa é:", p1.nome, p1.sobrenome, "\nIdade: ", p1.idade)
}

func fun01() func() {
	return func() {
		fmt.Println("Func 02")
	}
}

func teste01() {
	fmt.Println("Teste 01")
}

func chamaTeste01(x func()) {
	fmt.Println("Atenção:")
	x()
}

func a() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	fmt.Println("\n######## Exercício 01 ########")
	fmt.Println(exe01())
	fmt.Println("\n######## Exercício 02 ########")
	ex01 := []int{1, 2, 4, 5, 6, 80, 90, 100}
	fmt.Println(exe02(ex01...))
	defer fmt.Println("Será executado por último.")
	fmt.Println("\n######## Exercício 03 ########")
	fmt.Println("\n######## Exercício 04 ########")
	p2 := pessoa{
		nome:      "João",
		sobrenome: "Maria",
		idade:     40,
	}
	p2.exe04()
	fmt.Println("\n######## Exercício 05 ########")
	func() {
		fmt.Println("função anônima")
	}()
	fmt.Println("\n######## Exercício 06 ########")
	x := func() {
		fmt.Println("Atribuindo função anônima na variável")
	}
	x()
	fmt.Println("\n######## Exercício 07 ########")
	y := fun01()
	y()
	fmt.Println("\n######## Exercício 08 ########")
	chamaTeste01(teste01)
	fmt.Println("\n######## Exercício 09 ########")
	a1 := a()
	b1 := a()

	fmt.Println(a1())
	fmt.Println(a1())
	fmt.Println(a1())
	fmt.Println(a1())
	fmt.Println(a1())
	fmt.Println(b1())
	fmt.Println(b1())
	fmt.Println(b1())
	fmt.Println(b1())

}
