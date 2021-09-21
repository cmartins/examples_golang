package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracao bool
}

type sedan struct {
	veiculo
	modeluxo bool
}

func exe01() {
	pessoa01 := pessoa{
		nome:      "Cristiano",
		sobrenome: "Martins",
		sabores:   []string{"Chocolate", "Morango"},
	}
	fmt.Println(pessoa01.nome, pessoa01.sobrenome)
	for _, v := range pessoa01.sabores {
		fmt.Println("\t", v)
	}
}

func exe02() {
	carro_luxo := sedan{veiculo{4, "prata"}, true}
	caminhonete_4x4 := caminhonete{
		veiculo: veiculo{
			portas: 2,
			cor:    "vermelha",
		},
		tracao: true,
	}
	fmt.Println(carro_luxo)
	fmt.Println(caminhonete_4x4)
	fmt.Println(carro_luxo.cor)
	fmt.Println(caminhonete_4x4.tracao)
}

func exe03() {
	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Stroopwafel",
		sabor:   "Doce",
		ondetem: []string{"Holanda", "Shopping", "Na casa do tio rico"},
		vaibemcom: map[string]string{
			"Café da manhã": "Chá",
			"Almoço":        "Cafezinho"},
	}

	fmt.Println(x)
}

func main() {
	fmt.Println("\n######## Exercício 01 ########")
	exe01()
	fmt.Println("\n######## Exercício 02 ########")
	exe02()
	fmt.Println("\n######## Exercício 03 ########")
	exe03()

}
