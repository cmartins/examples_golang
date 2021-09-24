package main

import (
	"fmt"
)

func simplechannel() {
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func genSelect(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()
	return c
}

func receiveSelect(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func receive(canal <-chan int) {
	for v := range canal {
		fmt.Println(v)
	}
}

func exCommaOK() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func exechannel03(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func exechannel04() {
	canal := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, "\t", <-canal)
	}
}

func main() {
	fmt.Println("######## Channel 01 ########")
	simplechannel()
	fmt.Println("######## Channel 02 ########")
	c := gen()
	receive(c)
	fmt.Println("Exit")
	fmt.Println("######## Select statement ########")
	q := make(chan int)
	cs := genSelect(q)
	receiveSelect(cs, q)
	fmt.Println("Exit")
	fmt.Println("######## Ex comma ok ########")
	exCommaOK()
	fmt.Println("######## Channel 03 ########")
	cs03 := make(chan int)
	go exechannel03(cs03)
	for v := range cs03 {
		fmt.Println(v)
	}
	fmt.Println("######## Channel 04 ########")
	fmt.Println("valor: \t goroutine: ")
	exechannel04()
}
