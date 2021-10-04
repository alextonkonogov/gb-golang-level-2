package myErrors

import (
	"fmt"
)

func Example() {
	err := New("Alarm! An error has occured!")
	fmt.Println(err.Error())
	// Output:
	//Error text: Alarm! An error has occured! | Error time: 2021-10-04 11:14:31.1272 +0300 MSK m=+0.001379126
}
