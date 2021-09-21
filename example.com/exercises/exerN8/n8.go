package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type porIdade []user

func (p porIdade) Len() int           { return len(p) }
func (p porIdade) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p porIdade) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type porSobrenome []user

func (p porSobrenome) Len() int           { return len(p) }
func (p porSobrenome) Less(i, j int) bool { return p[i].Last < p[j].Last }
func (p porSobrenome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type jsontogo []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func exe01() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "James2",
		Age:   27,
	}

	u3 := user{
		First: "James3",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

}

func exe02() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var result jsontogo

	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		fmt.Println("deu ruim", err)
	}

	fmt.Println(result)
	fmt.Println(result[1])
	fmt.Println(result[1].Last)

}

func exe03() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("Deu ruim", err)
	}
}

func exe04() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)

	sort.Ints(xi)

	fmt.Println(xi)

	fmt.Println(xs)

	sort.Strings(xs)

	fmt.Println(xs)
}

func exe05() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(porIdade(users))
	fmt.Println("Ordenado por idade:\n")
	harmoniosa(users)

	sort.Sort(porSobrenome(users))
	fmt.Println("Ordenado por sobrenome:\n")
	harmoniosa(users)

}

func harmoniosa(x []user) {
	for i, v := range x {
		fmt.Println(i, "\tIdade:", v.Age, "\tNome completo:", v.First, v.Last, "\n")
		for _, c := range v.Sayings {
			fmt.Println("\t", c)
		}
		fmt.Println("\n")

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
}
