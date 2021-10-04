package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

//Для закрепления практических навыков программирования, напишите программу,
//которая создаёт один миллион пустых файлов в известной, пустой директории файловой системы используя вызов os.Create.
//Ввиду наличия определенных ограничений операционной системы на число открытых файлов, такая программа должна выполнять аварийную остановку.
//Запустите программу и дождитесь полученной ошибки.
//Используя отложенный вызов функции закрытия файла, стабилизируйте работу приложения.
//Критерием успешного выполнения программы является успешное создание миллиона пустых файлов в директории
func main() {
	dir := "oneMillionEmptyFilesFolder"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000000; i++ {
		err = myEmptyFileCreatingFunction(filepath.Join(dir, strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func myEmptyFileCreatingFunction(filename string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("from myEmptyFileCreatingFunction: %w", err)
	}
	defer file.Close()
	return
}
