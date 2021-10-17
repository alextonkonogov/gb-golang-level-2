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
	if true {
		go log.Printf("%s", "Four")
	}
	for i := 0; i < 3; i++ {
		go log.Println("Five")
	}
	switch true {
	case true:
		go log.Println("Six")
	}
	time.Sleep(5 * time.Second)
}
