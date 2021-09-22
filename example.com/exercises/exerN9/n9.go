package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "Hi!")
}

type humanos interface {
	falar()
}

func dizAlgumaCoisa(h humanos) {
	h.falar()
}

var wg sync.WaitGroup
var mu sync.Mutex
var contador int
var contadorAtomicEx int64

const quantidadeDeGoroutines = 20000

func exe01(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Goroutine: ", i+1)
			wg.Done()
		}(x)
	}
}

func exe02() {
	p1 := pessoa{"Maria", 30}
	(&p1).falar()
	dizAlgumaCoisa(&p1)
}

//go run -race n9.go
func exeMutexRaceProblem(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}

//go run -race n9.go
func exeMutexRaceFixed(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}

func exeAtomic(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt64(&contadorAtomicEx, 1)
			v := atomic.LoadInt64(&contadorAtomicEx)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}

func exeBuild() {
	fmt.Println("Seu OS é: \t", runtime.GOOS)
	fmt.Println("Seu ARCH é: \t", runtime.GOARCH)
}

func main() {
	fmt.Println("######## Exercício 01 ########")
	exe01(1000)
	wg.Wait()

	fmt.Println("######## Exercício 02 ########")
	exe02()
	fmt.Println("######## Exercício Mutex Race Problem ########")
	exeMutexRaceProblem(quantidadeDeGoroutines)
	wg.Wait()

	fmt.Println("Total de goroutines (Race problem):\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)
	fmt.Println("######## Exercício Mutex Race Problem fixed ########")
	exeMutexRaceFixed(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines: (Race fixed) \t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

	fmt.Println("######## Exercício Atomic ########")
	exeAtomic(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines atomic:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contadorAtomicEx)

	fmt.Println("######## Exercício Build ########")
	exeBuild()

}
