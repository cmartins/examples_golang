package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func testError01() {
	p1 := person{
		First:   "Maria",
		Last:    "Santos",
		Sayings: []string{"Ok", "Not ok"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}

func testError02() {
	p1 := person{
		First:   "Maria",
		Last:    "Santos",
		Sayings: []string{"Ok", "Not ok"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("Deu ruim")
	}
	fmt.Println(string(bs))
}

type erroCustomizado struct {
	msgerro string
}

func (e erroCustomizado) Error() string {
	return fmt.Sprintf("Deu ruim, olha o erro: %v", e.msgerro)
}

func erroComoParametro(e error) {
	fmt.Println("Passaram como argumento: ", e.(erroCustomizado).msgerro, "\nno método do Error, temos o erro:", e)
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("deu ruim")
	}
	return bs, nil
}

func main() {
	fmt.Println("######## Teste 01 ########")
	testError01()
	fmt.Println("######## Teste 02 ########")
	testError02()
	fmt.Println("######## Teste 03 ########")
	erroArg := erroCustomizado{"Teste!!!"}
	erroComoParametro(erroArg)
}
