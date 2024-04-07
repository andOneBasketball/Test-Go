package main

import "fmt"

type FileMd5Info struct {
	Product            string
	Devid              string
}

func main() {
    md5DataCache := make([]FileMd5Info, 0, 10)
	a := FileMd5Info{Product: "a", Devid: "1234aef"}
	md5DataCache = append(md5DataCache, a)
	b := FileMd5Info{Product: "b", Devid: "14aef"}
	md5DataCache = append(md5DataCache, b)
	c := FileMd5Info{Product: "c", Devid: "14a12ef"}
	md5DataCache = append(md5DataCache, c)
	d := FileMd5Info{Product: "d", Devid: "112ef"}
	md5DataCache = append(md5DataCache, d)

	fmt.Printf("%#v\n", md5DataCache)

	left := md5DataCache[:2]
	right := md5DataCache[3:]
	md5DataCache = append(left, right...)
	fmt.Printf("%#v\n", md5DataCache)
}

