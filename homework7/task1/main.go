package main

import (
	"fmt"
	"log"

	"github.com/alextonkonogov/gb-golang-level-2/homework7/task1/funcs"
	"github.com/alextonkonogov/gb-golang-level-2/homework7/task1/persons"
)

//Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и values map[string]interface{}
//(key - название поля структуры, которому нужно присвоить value этой мапы).
//Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
//Функция может возвращать только ошибку error. Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).
func main() {
	m := map[string]interface{}{
		"Name":        "Alex",
		"Age":         33,
		"Married":     true,
		"Temperature": 36.6,
	}

	person := persons.Person{}
	err := funcs.ChangeStructField(&person, m)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(person)
}
