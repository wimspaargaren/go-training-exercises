package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

type realFileHandler struct{}

func (r realFileHandler) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func NewRealFileHandler() fs.FS {
	return realFileHandler{}
}

func main() {
	realFileHandler := NewRealFileHandler()
	ioHandler := NewIOHandler(realFileHandler)
	fileContent, err := ioHandler.readFile("main.go")
	if err != nil {
		log.Fatalf("unable to read file: %s", err)
	}
	fmt.Println("file content", fileContent)
}

type ioHandler struct {
	fs fs.FS
}

func NewIOHandler(fs fs.FS) ioHandler {
	return ioHandler{
		fs: fs,
	}
}

// Write a unit test for readFile by creating a custom implementation for fs.FS interface in ioHandler.
func (i ioHandler) readFile(fileName string) (string, error) {
	f, err := i.fs.Open(fileName)
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
