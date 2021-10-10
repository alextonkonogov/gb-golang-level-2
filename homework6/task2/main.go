package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

//Написать многопоточную программу, в которой будет использоваться явный вызов планировщика. Выполните трассировку программы

//Если закомментировать runtime.Gosched(), убрать трассировку и запустить на одном ядре (GOMAXPROCS=1), то программа выводет 5 раз только "hello"
//В текущем состоянии в консоль поочередно будет выводиться "hello" и "world". Принудительный вызов планировщика помогает получить время выполнения другим потокам программы
//см. https://drive.google.com/drive/folders/1b0rDMIM2cL0gpm-qtAcXmCwm1_zHgCt3?usp=sharing
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	go say("world")
	say("hello")
}
