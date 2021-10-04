package myErrors

import (
	"fmt"
	"time"
)

type myError struct {
	text string
	time time.Time
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error text: %s | Error time: %v\n", e.text, e.time)
}

func New(text string) *myError {
	return &myError{
		text: text,
		time: time.Now(),
	}
}
