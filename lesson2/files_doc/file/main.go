// Package implements and use function for ceating new files
// The number of files that will be created you set yourself
//
// Function accepts an integer argument
//
// createFile(n int)
//
// And creates such number of files that you sent as it's argument
package file

import (
	"fmt"
	"os"
)

// CreateFile creates files up in file structure by path "../../data"
// Files will have names starting wit "new_file" and ending it's index
// number
func CreateFile(n int) {
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