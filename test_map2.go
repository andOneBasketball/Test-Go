package main

import (
	"fmt"
)

// 文件指纹信息数据结构
type IdentifyEng struct {
	dropWhite           bool
	blackFingerprint    map[string]map[string]bool
}

func f1(s *map[string]map[string]bool) {
	fmt.Printf("%#v %#v\n", *s, *s["abcd"]["1234"])
}

func main() {
	a := IdentifyEng{}
	a.blackFingerprint = make(map[string]map[string]bool)
	a.blackFingerprint["abcd"] = make(map[string]bool)
	if _, valueExist := a.blackFingerprint["abcd"]; valueExist == true {
		a.blackFingerprint["abcd"]["1234"] = true
	}
	fmt.Printf("%#v\n", a.blackFingerprint)
}
