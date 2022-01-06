package main

import (
	"fmt"
	"time"
)

func selfRecover() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered after", v)
		}
	}()
	panic("A-A-A!!!")
}

func main() {
	go selfRecover()
	time.Sleep(time.Second)
}
