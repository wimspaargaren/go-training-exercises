package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

type realFileHandler struct {
}

func (r realFileHandler) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func NewRealFileHandler() fs.FS {
	return realFileHandler{}
}

func main() {
	realFileHandler := NewRealFileHandler()
	fileContent, err := readFile(realFileHandler)
	if err != nil {
		log.Fatalf("unable to read file: %s", err)
	}
	fmt.Println("file content", fileContent)
}

// Write a unit test for readFile by creating a custom implementation for fs.FS interface.
func readFile(fileSystem fs.FS) (string, error) {
	f, err := fileSystem.Open("main.go")
	if err != nil {
		return "", err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println("couldn't close file:", err)
		}
	}()
	res, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
