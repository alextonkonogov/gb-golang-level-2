package main

import (
	"fmt"
	"time"
)

//Выполните задание из блока “Для самостоятельного изучения” данной методички
//Для самостоятельного изучения. Предложите реализацию примера так, чтобы аварийная остановка программы не выполнилась.
func main() {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered", v)
			}
		}()

		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}
