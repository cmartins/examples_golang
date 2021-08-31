package main

import (
	"fmt"
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

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)

}

func main() {
	exer01()
	exer02()
}
