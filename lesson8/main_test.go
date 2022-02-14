package main_test

import (
	"errors"
	. "github.com/golang2/lesson8"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"testing"
)

var hashes = map[string]string{
	"./test/4.txt": "4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a",
}

func TestIsDirectoryIfDir(t *testing.T) {
	result, err := IsDirectory("./test/1")
	if err != nil {
		t.Errorf("Not expected error")
	}
	if result == false {
		t.Errorf("Should be a directory")
	}
}

func TestIsDirectoryIfFile(t *testing.T) {
	result, err := IsDirectory("./test/4.txt")
	if err != nil {
		t.Errorf("Not expected error")
	}
	if result == true {
		t.Errorf("Should be a file")
	}
}

func TestInnerWorkerIfFile(t *testing.T) {
	path := "./test/4.txt"
	dirPath, fileHash, err := InnerWorker(path)
	if err != nil {
		t.Errorf("Not expected error")
	}
	if len(dirPath) != 0 || fileHash == nil {
		t.Errorf("Should be a file")
	}
	fullPath, _ := filepath.Abs(path)
	if fileHash[fullPath] != hashes[path] {
		t.Errorf("Incorrect hash")
	}
}

func TestInnerWorkerIfDir(t *testing.T) {
	path := "./test/1"
	dirPath, fileHash, err := InnerWorker(path)
	if err != nil {
		t.Errorf("Not expected error")
	}
	if fileHash != nil {
		t.Errorf("Should be a file")
	}
	if len(dirPath) != 10 {
		t.Errorf("Incorrect calculated number of child dirs")
	}
}

func TestCopyHashesWithEmptyArguments(t *testing.T) {
	allHashes := make(FileToHashMap)
	fHashes := make(FileToHashMap)
	CopyHashes(&allHashes, &fHashes)
	if len(allHashes) != 0 {
		t.Errorf("Should be empty")
	}
}

func TestCopyHashesCopyToTarget(t *testing.T) {
	allHashes := make(FileToHashMap)
	fHashes := FileToHashMap{
		"key": "value",
	}
	CopyHashes(&allHashes, &fHashes)
	if len(allHashes) != 1 {
		t.Errorf("Wrong copying")
	}
	if allHashes["key"] != "value" {
		t.Errorf("Wrong copying")
	}
}

func TestCopyHashesNotCopyToSource(t *testing.T) {
	allHashes := FileToHashMap{
		"key": "value",
	}
	fHashes := make(FileToHashMap)
	CopyHashes(&allHashes, &fHashes)
	if len(allHashes) != 1 {
		t.Errorf("Wrong copying")
	}
	if allHashes["key"] != "value" {
		t.Errorf("Wrong copying")
	}
	if len(fHashes) == 1 {
		t.Errorf("Wrong copying")
	}
}

func TestRemoveDuplicatesIfEmptyArgument(t *testing.T) {
	allHashes := make(FileToHashMap)
	err := RemoveDuplicates(allHashes)
	if err != nil {
		t.Errorf("Not expected error")
	}
}

func TestRemoveDuplicatesIfOnlyOneFile(t *testing.T) {
	allHashes := FileToHashMap{
		"key": "value",
	}
	err := RemoveDuplicates(allHashes)
	if err != nil {
		t.Errorf("Not expected error")
	}
}

func TestRemoveDuplicatesIfCopyWasNotFound(t *testing.T) {
	allHashes := FileToHashMap{
		"key": "value",
		"key2": "value",
	}
	err := RemoveDuplicates(allHashes)
	if err == nil {
		t.Errorf("Should be error because key file was not found")
	}
}

func TestRemoveDuplicatesIfCopy(t *testing.T) {
	defer func (){
		src := "./test/4.txt"
		dest := "./test/1/6/4.txt"

		bytesRead, err := ioutil.ReadFile(src)

		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(dest, bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
		}
	}()
	allHashes := FileToHashMap{
		"./test/4.txt": "value",
		"./test/1/6/4.txt": "value",
	}
	err := RemoveDuplicates(allHashes)
	if err != nil {
		t.Errorf("Unexpected error")
	}
	if _, err := os.Stat("./test/1/6/4.txt"); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("Unsuccessfull delete")
	}
}
