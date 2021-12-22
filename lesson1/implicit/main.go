package main

import (
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"
)

type ErrorWithTimeAndTrace struct {
	text  string
	time  string
	trace string
}

func New(text string) error {
	t := time.Now()
	return &ErrorWithTimeAndTrace{
		text:  text,
		time:  t.String(),
		trace: string(debug.Stack()),
	}
}

func (e *ErrorWithTimeAndTrace) Error() string {
	return fmt.Sprintf("error: %s\ntime: %s\ntrace:\n%s", e.text, e.time, e.trace)
}

func getValue(index int) (result int, err error) {
	slice := []int{}

	defer func() {
		if v := recover(); v != nil {
			err = New("Index out of range")
		}
	}()

	return slice[index], err
}

func main() {
	result, err := getValue(rand.Int())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
