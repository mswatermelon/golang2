package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

const goroutinesCount = 999

type Storage struct {
	sync.Mutex
	greetingStr         string
	greetingBytes       []byte
	isConvertingToBytes bool
}

func (s *Storage) convertToBytes() {
	s.greetingBytes = []byte(s.greetingStr)
	s.greetingStr = ""
}

func (s *Storage) convertToString() {
	s.greetingStr = string(s.greetingBytes)
	s.greetingBytes = nil
}

func (s *Storage) convert() {
	s.Lock()
	defer s.Unlock()
	if s.isConvertingToBytes {
		s.convertToBytes()
	} else {
		s.convertToString()
	}
	s.isConvertingToBytes = !s.isConvertingToBytes
}

func (s *Storage) printStrValue() {
	fmt.Println(s.greetingStr)
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	converterStorage := Storage{
		greetingStr:         "Hello, dear friend",
		isConvertingToBytes: true,
	}

	for i := 0; i < goroutinesCount; i++ {
		go converterStorage.convert()
	}

	converterStorage.printStrValue()
}
