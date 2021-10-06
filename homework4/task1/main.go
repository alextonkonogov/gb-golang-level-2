package main

import (
	"fmt"
	"sync"
)

//С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из которых увеличивает число на 1.
//Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000.
func main() {
	var counter int
	n := 1000
	ch := make(chan int, n)
	wg := sync.WaitGroup{}

	wg.Add(n)
	go func(ch <-chan int) {
		for v := range ch {
			counter += v
		}
		wg.Done()
	}(ch)

	for i := 0; i < n; i++ {
		go func(ch chan<- int) {
			ch <- 1
			wg.Done()
		}(ch)
	}

	wg.Wait()
	fmt.Printf("The counter is: %v\n", counter)
}
