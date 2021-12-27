package main

import (
	"fmt"
	"os"
)

func createFile(n int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	f, err := os.Create(fmt.Sprint("../../data/new_file", n))
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func main() {
	for i := 0; i < 1000000; i++ {
		createFile(i)
	}
}
