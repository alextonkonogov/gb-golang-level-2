package main

import (
	"fmt"
	"log"
	"time"
)

func MyTestNew() {
	go log.Println("One")
	go fmt.Println("Two")
	time.Sleep(5 * time.Second)
}

func TheFunctionWithSomeAsyncFunctionsInside() {
	go log.Println("One")
	go fmt.Println("Two")
	go log.Printf("%s", "Three")
	time.Sleep(5 * time.Second)
}
