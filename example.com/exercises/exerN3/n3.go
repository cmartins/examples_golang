package main

import (
	"fmt"
)

func exe01() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exe02() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%U\n", i)
		}
	}
}

func exe03() {
	birthdate := 1983
	ano_count := 2088

	for birthdate <= ano_count {
		fmt.Println(birthdate)
		birthdate++
	}
}

func exe04() {
	birthdate := 1983
	ano_count := 2088

	for {
		if birthdate > ano_count {
			break
		}
		fmt.Println(birthdate)
		birthdate++
	}
}

func exe05() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

func exe06() {
	is_ok := true
	if is_ok {
		fmt.Println("Ok")
	}
}

func exe07() {
	vl := 2
	if vl == 0 {
		fmt.Println("A")
	} else if vl == 1 {
		fmt.Println("B")
	} else {
		fmt.Println("...")
	}
}

func exe08() {
	vl := 2

	switch {
	case vl == 0:
		fmt.Println("A")
	case vl == 1:
		fmt.Println("B")
	case vl == 2:
		fmt.Println("C")
	}
}

func exe09() {
	vl := 2

	switch vl {
	case 0:
		fmt.Println("A")
	case 1:
		fmt.Println("B")
	case 2:
		fmt.Println("C")
	}
}

func main() {
	fmt.Println("######## Exercício 01 ########")
	exe01()
	fmt.Println("######## Exercício 02 ########")
	exe02()
	fmt.Println("######## Exercício 03 ########")
	exe03()
	fmt.Println("######## Exercício 04 ########")
	exe04()
	fmt.Println("######## Exercício 05 ########")
	exe05()
	fmt.Println("######## Exercício 06 ########")
	exe06()
	fmt.Println("######## Exercício 07 ########")
	exe07()
	fmt.Println("######## Exercício 08 ########")
	exe08()
	fmt.Println("######## Exercício 09 ########")
	exe09()
}
