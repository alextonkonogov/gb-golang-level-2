package main

import (
	"fmt"
	"github.com/alextonkonogov/gb-golang-level-2/homework1/task2/myErrors"
	"log"
)

//Дополните функцию из п.1 возвратом собственной ошибки в случае возникновения панической ситуации.
//Собственная ошибка должна хранить время обнаружения панической ситуации.
//Критерием успешного выполнения задания является наличие обработки созданной ошибки в функции main и вывод ее состояния в консоль
func main() {
	err := myFuncWithRecoveredPanic()
	if err != nil {
		log.Println(err)
	}
}

func myFuncWithRecoveredPanic() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = myErrors.New(fmt.Sprintf("recovered %v", v))
		}
	}()

	arr := []uint{0, 1, 2, 3}
	for i := 0; i <= 5; i++ {
		fmt.Printf("index: %d, value: %d\n", i, arr[i])
	}

	return
}
