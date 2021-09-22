package main

import (
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
	os.MkdirAll(dir, os.ModePerm)
	for i := 0; i < 1000000; i++ {
		err := myEmptyFileCreatingFunction(filepath.Join(dir, strconv.Itoa(i)))
		if err != nil {
			panic(err)
		}
	}
}

func myEmptyFileCreatingFunction(filename string) (err error) {
	file, err := os.Create(filename)
	defer file.Close()
	return
}
