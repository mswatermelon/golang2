package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			runtime.Gosched()
		}
		fmt.Println(s)
	}
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	go say("world")
	say("hello")
}
