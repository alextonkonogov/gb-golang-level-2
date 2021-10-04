package main

import (
	"fmt"
	"github.com/alextonkonogov/gb-golang-level-2/homework2/tasks1and2/myErrors"
	"log"
)

//task 1
//Выполните сборку ваших предыдущих программ под операционную систему, отличающуюся от текущей.
//Проанализируйте вывод команды file для полученного исполняемого файла.
//Попробуйте запустить исполняемый файл

//task 2
//Напишите документацию для одной из предыдущих программ. Запустите сервер документации локально.
//Убедитесь, что документация отображается корректно.

//На всякий приложил скрины: https://drive.google.com/drive/folders/1lf9J5WV_BkXz9uNHhIY8u0Hat_VehnRV?usp=sharing
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
