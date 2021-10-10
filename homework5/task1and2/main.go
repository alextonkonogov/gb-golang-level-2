package main

import (
	"flag"
	"fmt"
	"sync"
)

//Task 1 Напишите программу, которая запускает n потоков и дожидается завершения их всех
//Task 2 Реализуйте функцию для разблокировки мьютекса с помощью defer

var amount = flag.Int("amount", 10, "amount of goroutines")

func main() {
	flag.Parse()
	fmt.Printf("Program is working...\n__amount of set goroutines = %v\n", *amount)

	var finished = struct {
		sync.Mutex
		count int
	}{}
	wg := sync.WaitGroup{}
	wg.Add(*amount)

	for i := 0; i < *amount; i++ {
		go func() {
			finished.Lock()
			defer finished.Unlock()
			finished.count += 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("__amount of finished goroutines = %v\nProgram has finished\n", finished.count)
}
