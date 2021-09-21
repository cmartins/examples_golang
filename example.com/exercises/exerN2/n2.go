package main

import (
	"fmt"
)

const exe03_x int = 10
const exe03_y = 10
const (
	_ = 2021 + iota
	exe06_b
	exe06_c
	exe06_d
	exe06_e
)

func exer01() {
	x := 100
	fmt.Printf("%d, %#x, %b", x, x, x)
}

func exer02() {
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 >= 100)
	f := (10 > 100)

	fmt.Printf("\n%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}

func exer03() {
	fmt.Printf("\nx: %v - y: %v", exe03_x, exe03_y)
}

func exer04() {
	x := 300
	fmt.Printf("\n%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)
}

func exer05() {
	x := `
		Algum texto
					com 

					formatacao
	`
	fmt.Println(x)
}

func exer06() {
	fmt.Println(exe06_b, exe06_c, exe06_d, exe06_e)
}

func main() {
	fmt.Println("######## Exercicio 01 ########")
	exer01()
	fmt.Println("\n\n######## Exercicio 02 ########")
	exer02()
	fmt.Println("\n\n######## Exercicio 03 ########")
	exer03()
	fmt.Println("\n\n######## Exercicio 04 ########")
	exer04()
	fmt.Println("\n\n######## Exercicio 05 ########")
	exer05()
	fmt.Println("\n\n######## Exercicio 06 ########")
	exer06()
}
