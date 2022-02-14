package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	targetPath = flag.String("path", ".", "Target directory path")
	workersCount = flag.Int("count", runtime.NumCPU(), "Count of workers that will process directory")
	force = flag.Bool("force", true, "if force - continue on error")
)

type FileToHashMap = map[string]string

type DirData struct {
	directories []string
	fileHash    FileToHashMap
	err         error
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func innerWorker(source string) (dirPath []string, fileHash FileToHashMap, err error) {
	path, err := filepath.Abs(source)
	if err != nil {
		return dirPath, fileHash, err
	}
	isDir, err := isDirectory(path)
	if err != nil {
		return dirPath, fileHash, err
	}
	if isDir {
		items, err := ioutil.ReadDir(path)
		if err != nil {
			return dirPath, fileHash, err
		}
		result := make([]string, len(items))
		for i := range items {
			result[i] = filepath.Join(path, items[i].Name())
		}
		return result, nil, nil
	}
	h := sha256.New()
	f, err := os.Open(path)
	if err != nil {
		return dirPath, fileHash, err
	}
	if _, err := io.Copy(h, f); err != nil {
		if err := f.Close(); err != nil {
			return dirPath, fileHash, err
		}
		return dirPath, fileHash, err
	}
	fileHash = FileToHashMap{
		path: hex.EncodeToString(h.Sum(nil)),
	}
	if err := f.Close(); err != nil {
		return dirPath, fileHash, err
	}
	return nil, fileHash, nil
}

func worker(
	children []string,
	directories chan DirData,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for _, dir := range children {
		dirItems, fileHash, err := innerWorker(dir)
		if err != nil {
			if *force {
				continue
			}
			break
		}
		directories <- DirData{
			directories: dirItems,
			fileHash:    fileHash,
			err:         err,
		}
	}
}

func investigate(workersCount int, children []string) chan DirData {
	var wg sync.WaitGroup

	childrenLen := len(children)
	directories := make(chan DirData)
	tasksCount := int(math.Max(float64(childrenLen/workersCount), 1))
	x := 0
	if workersCount > tasksCount {
		workersCount = tasksCount
	}
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		if i == workersCount - 1 {
			go worker(children[x:childrenLen], directories, &wg)
		} else {
			go worker(children[x:x+tasksCount], directories, &wg)
			x += tasksCount
		}
	}

	go func() {
		wg.Wait()
		close(directories)
	}()

	return directories
}

func copyHashes(targetMap *FileToHashMap, originalMap *FileToHashMap) {
	for key, value := range *originalMap {
		(*targetMap)[key] = value
	}
}

func main()  {
	flag.Parse()
	children := make([]string, 0)
	if len(*targetPath) == 0 {
		err := fmt.Errorf("Target directory path can not be empty")
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	children = append(children, *targetPath)

	allHashes := make(FileToHashMap)
	for len(children) > 0 {
		var newChildren []string
		fHashes := make(FileToHashMap)

		directories := investigate(*workersCount, children)

		for i := range directories {
			if i.err != nil {
				fmt.Printf(i.err.Error())
				os.Exit(1)
			}
			if i.fileHash == nil {
				newChildren = append(newChildren, i.directories...)
			} else {
				hash := i.fileHash
				if hash != nil {
					for k, v := range hash {
						fHashes[k] = v
					}
				}
			}
		}
		copyHashes(&allHashes, &fHashes)
		children = newChildren
	}

	hashToFilesMap := make(map[string][]string)
	for key, value := range allHashes {
		hashToFilesMap[value] = append(hashToFilesMap[value], key)
	}
	for _, value := range hashToFilesMap {
		if len(value) <= 1 {
			continue
		}
		for _, file := range value[1:] {
			fmt.Println(file)
			err := os.Remove(file)

			if err != nil {
				fmt.Printf(err.Error())
				os.Exit(1)
			}
		}
		fmt.Println()
	}
}