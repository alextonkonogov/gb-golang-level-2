package main

import (
	"fmt"
	"sync"
)

//С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из которых увеличивает число на 1.
//Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000.
func main() {
	var counter = struct {
		sync.Mutex
		n int
	}{}
	workers := 1000

	ch := make(chan int, workers)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(ch chan<- int) {
			counter.Lock()
			defer counter.Unlock()
			counter.n += 1
			wg.Done()
		}(ch)
	}

	wg.Wait()
	fmt.Printf("The counter is: %v\n", counter.n)
}
