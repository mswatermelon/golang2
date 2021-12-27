package main

import (
	file "github.com/mswatermelon/lesson2/files/file"
)

func main() {
	for i := 0; i < 1000000; i++ {
		file.CreateFile(i)
	}
}
