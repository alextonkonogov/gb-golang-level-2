package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее,
//чем за одну секунду (установить таймаут).
//скрин с результатом: https://drive.google.com/file/d/1FCeRFhhQGjHOGQQYxMulnJ0xPvB7xIKC/view?usp=sharing

var timeOut = flag.Int64("timeout", 1, "context time out")

func main() {
	flag.Parse()
	fmt.Printf("Program is working...\n    timeout = %v sec.\n    execute \"kill -SIGTERM $(pgrep main)\" to see the result\n", *timeOut)
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGTERM)
	exit_chan := make(chan int)

	go func() {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(*timeOut)*time.Second)
		defer cancelFunc()

		for {
			s := <-signalChanel
			switch s {
			case syscall.SIGTERM:
				select {
				case <-ctx.Done():
					fmt.Println(ctx.Err())
					exit_chan <- 1
				default:
					fmt.Println("SIGTERM was received")
					exit_chan <- 0
					fmt.Println("...program has been stopped")
				}
			default:
				fmt.Println("Unknown signal")
				exit_chan <- 1
			}
		}
	}()
	exitCode := <-exit_chan
	os.Exit(exitCode)
}
