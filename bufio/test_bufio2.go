package main

// 统计不同文件出现相同内容的行数

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if fmap, ok1 := counts[line]; ok1 {
				if _, ok2 := fmap[filename]; ok2 {
					fmap[filename]++
					counts[line] = fmap
				} else {
					fmap[filename] = 1
					counts[line] = fmap
				}
			} else {
				counts[line] = map[string]int{filename: 1}
			}
		}
	}
	//fmt.Printf("counts: %v\n", counts)
	
	for line, fmap := range counts {
		for filename, count := range fmap {
			fmt.Printf("filenmae: %s, line: %s, count: %d\n", filename, line, count)
		}
	}
}