// Package myErrors implements implements the error interface a adds to it time when errors happened.
// To create an error use function New and then you can get a string with placeholded time through method Error()
package myErrors

import (
	"fmt"
	"time"
)

type myError struct {
	text string
	time time.Time
}

// The Error method returns a string with a text and time when the error has happened:
// fmt.Sprintf("Error text: %s | Error time: %v", e.text, e.time)
func (e *myError) Error() string {
	return fmt.Sprintf("Error text: %s | Error time: %v", e.text, e.time)
}

//New function creates error with a text message and a time inside,
// gets a text of an error as the only argument, authomatically gets current time and returns err.
func New(text string) error {
	return &myError{
		text: text,
		time: time.Now(),
	}
}
