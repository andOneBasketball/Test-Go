package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
	"strings"
	"path/filepath"
)

func Unzip(zipath, dir string) error {
	// Open zip file.
	reader, err := zip.OpenReader(zipath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := unzipFile(file, dir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dir string) error {
	// Prevent path traversal vulnerability.
	// Such as if the file name is "../../../path/to/file.txt" which will be cleaned to "path/to/file.txt".
	name := strings.TrimPrefix(filepath.Join(string(filepath.Separator), file.Name), string(filepath.Separator))
	filePath := filepath.Join(dir, name)

	// Create the directory of file.
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Open the file.
	r, err := file.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	// Create the file.
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// Save the decompressed file content.
	_, err = io.Copy(w, r)
	return err
}


func main() {
    // 注册解压函数
    zip.RegisterDecompressor(zip.Store, func(in io.Reader) io.ReadCloser {
        return &passwordReader{in, []byte("abcd1234")}
    })

    // 打开 ZIP 文件
    r, err := zip.OpenReader("./test.zip")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer r.Close()

    // 解压文件
    for _, f := range r.File {
        fmt.Printf("Extracting %s\n", f.Name)
        err := extractFile(f)
        if err != nil {
            fmt.Println(err)
        }
    }
}

type passwordReader struct {
    r    io.Reader
    pass []byte
}

func (pr *passwordReader) Read(p []byte) (n int, err error) {
    n, err = pr.r.Read(p)
    for i := 0; i < n; i++ {
        p[i] ^= pr.pass[i%len(pr.pass)]
    }
    return
}

func (pr *passwordReader) Close() error {
    if c, ok := pr.r.(io.Closer); ok {
        return c.Close()
    }
    return nil
}

func extractFile(f *zip.File) error {
    // 打开 ZIP 文件中的文件
    rc, err := f.Open()
    if err != nil {
        return err
    }
    defer rc.Close()

    // 创建目标文件
    dst, err := os.Create(f.Name)
    if err != nil {
        return err
    }
    defer dst.Close()

    // 复制文件内容
    _, err = io.Copy(dst, rc)
    if err != nil {
        return err
    }

    return nil
}