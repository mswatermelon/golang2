package main

import (
	"fmt"
	"time"
)

type Handler func()

func createGorutine() {
	panic("A-A-A!!!")
}

func runGorutine() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	panic("A-A-A!!!")
}

func main() {
	go runGorutine()
	time.Sleep(time.Second)
}
