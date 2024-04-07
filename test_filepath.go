package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
)

func main() {
	path := "path/to/file.txt"
	filename := filepath.Base(path)
	fmt.Println(filename)


	dir := "path/to"
	filename = "file.txt"
	path = filepath.Join(dir, filename, "abcd.txt")
	fmt.Println(path, reflect.TypeOf(path))

	filename = "/a/we/w/example.123.txt"
    ext := strings.TrimPrefix(filepath.Ext(filename), ".")
    fmt.Println(ext)
}