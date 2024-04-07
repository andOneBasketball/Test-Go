package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"path/filepath"
)

func createFile(name string) (*os.File, error) {
    err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	fmt.Printf("name: %s, dirname: %s\n", name, string([]rune(name)[0:strings.LastIndex(name, "/")]))
    if err != nil {
        return nil, err
    }
    return os.Create(name)
}

func tarGzUnzip(zipFile, dest string) (error, []string) {
	var subFilenames []string
	fr, err := os.Open(zipFile)
	if err != nil {
		return err, subFilenames
	}
	defer fr.Close()
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err, subFilenames
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err, subFilenames
		}
		filename := filepath.Join(dest, h.Name)
		fmt.Printf("filename: %s\n", filename)
		filename = dest + h.Name
		fw, err := createFile(filename)
        if err != nil {
            return err, subFilenames
        }
		_, err = io.Copy(fw, tr)
		if err != nil {
			fw.Close()
			return err, subFilenames
		}
		fw.Close()
		subFilenames = append(subFilenames, h.Name)
	}
	return nil, subFilenames
}

func main() {
	subFilename := filepath.Base("./tmp/test.tar.gz")
	subFilepath := filepath.Join("./tmp", strings.Replace(subFilename, ".", "_", -1))
	err, subFilenames := tarGzUnzip("./tmp/test.tar.gz", "./tmp/test_tar_gz")
	//fmt.Printf("%v %v\n", err, subFilenames)
	fmt.Printf("%v %v %s %s\n", err, subFilenames, filepath.Join("./tmp/test_tar_gz", subFilenames[0]), subFilepath)
}
