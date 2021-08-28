package main

import "fmt"

var x int
var y string
var z bool

type hotdog int

var dog hotdog
var y_exe_05 int

func exe01() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(x, y, z)
}

func exe02() {
	fmt.Println(x, y, z)
}

func exe03() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("x: %v\t y: %v\t z: %v\t", x, y, z)
	fmt.Println(s)
}

func exe04() {
	fmt.Println(dog)
	fmt.Printf("%T\n", x)
	dog = 42
	fmt.Println(dog)
}

func exe05() {
	fmt.Printf("%v\n", dog)
	fmt.Printf("%T\n", dog)
	dog = 43
	fmt.Printf("%v\n", dog)
	y_exe_05 = int(dog)
	fmt.Println(y_exe_05)
	fmt.Printf("%T\n", y_exe_05)
}

func main() {
	fmt.Println("\nExercicio 01")
	exe01()
	fmt.Println("\nExercicio 02")
	exe02()
	fmt.Println("\nExercicio 03")
	exe03()
	fmt.Println("\nExercicio 04")
	exe04()
	fmt.Println("\nExercicio 05")
	exe05()
}
