package tool

import (
    "archive/tar"
    "compress/gzip"
    "io"
    "os"
	"strings"
	"path/filepath"
)

/*
@brief 创建文件，当目录不存在时先创建目录
@param[in] name 待创建的文件
@return 成功时返回创建成功的文件句柄和 nil
*/
func createFile(name string) (*os.File, error) {
    err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
    if err != nil {
        return nil, err
    }
    return os.Create(name)
}

/*
@brief 解压 tar gzip 包
@param[in] zipFile 压缩包
@param[in] dest 解压后文件的存放目录
@return 解压成功后返回 nil 和解压后的文件列表
*/
func DecompressTarGz(zipFile, dest string) (error, []string) {
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