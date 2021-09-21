package main

import (
	"fmt"
)

func exe01() {
	array := [5]int{1, 2, 3, 4, 5}

	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", array)
}

func exe02() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slice)
}

func exe03() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}

func exe04() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exe05() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func exe06() {
	x := make([]string, 26, 26)
	x = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(x), cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}

func exe07() {
	ss := [][]string{
		[]string{
			"A",
			"B",
			"C",
		},
		[]string{
			"D",
			"E",
			"F",
		},
		[]string{
			"G",
			"H",
			"I",
		},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Println("\n\n")

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}

}

func exe08() {
	mp := map[string][]string{
		"A": []string{
			"1A", "2A",
		},
		"B": []string{
			"1B", "2B",
		},
		"C": []string{
			"1C", "2C",
		},
	}

	for k, v := range mp {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}

func exe09() {
	mp := map[string][]string{
		"A": []string{
			"1A", "2A",
		},
		"B": []string{
			"1B", "2B",
		},
		"C": []string{
			"1C", "2C",
		},
	}

	mp["D"] = []string{"1D", "2D"}

	for k, v := range mp {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}

func main() {
	fmt.Println("######## Exercício 01 ########")
	exe01()
	fmt.Println("\n######## Exercício 02 ########")
	exe02()
	fmt.Println("\n######## Exercício 03 ########")
	exe03()
	fmt.Println("\n######## Exercício 04 ########")
	exe04()
	fmt.Println("\n######## Exercício 05 ########")
	exe05()
	fmt.Println("\n######## Exercício 06 ########")
	exe06()
	fmt.Println("\n######## Exercício 07 ########")
	exe07()
	fmt.Println("\n######## Exercício 08 ########")
	exe08()
	fmt.Println("\n######## Exercício 09 ########")
	exe09()
}
