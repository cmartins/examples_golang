package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func exer01() {
	x := 10
	fmt.Println(&x)
}

func exer02(p *pessoa) {
	(*p).nome = "Maria"
	p.sobrenome = "Joana"
}

func main() {
	fmt.Println("######## Exercicio 01 ########")
	exer01()
	fmt.Println("######## Exercicio 02 ########")
	jl := pessoa{"Jl", "Prima", 10}
	fmt.Println(jl)
	exer02(&jl)
	fmt.Println(jl)
}
