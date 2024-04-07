package main
 
import (
	"fmt"
	"strings"
)

const Alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func SoapaEscape(pattern string) string {
	var result = make([]string, len(pattern))

	for i, v := range pattern {
		result[i] = string(v)
	}

	for index, ch := range pattern {
		if false == strings.ContainsRune(Alphanum, ch) {
			if ch == '\000' {
				result[index] = "\\000"
			} else {
				result[index] = "\\" + string(ch)
			}
		}
	}
	return strings.Join(result, pattern[:0])
}
 
//golang字符串操作
func main(){
	s := "Hello world hello world"
	str := "Hello"
	//var s = []string{"11","22","33"}

	//返回s的副本，并将副本中的old字符串替换为new字符串，替换次数为n次，
	//如果n为-1，则全部替换；如果 old 为空，则在副本的每个字符之间都插入一个new
	ret := strings.Replace(s,"hello",str,1)
	fmt.Println(ret) //  Hello world Hello world

	a := "\u0000\u0000\u0000\u0002md5_check\u0000\r\u0000\u0000\u0000200, success\u0000\u0002log_check\u0000\r\u0000\u0000\u0000200, success\u0000\u0000"
	fmt.Println(string(a))
	b := "\r\u0000\u0000\u0000200, success\u0000"
	fmt.Println(string(b))

	s = "/bin/few你好newe,这里"
	fmt.Println("^" + SoapaEscape(s))

	s = "hello world"
    if ! strings.Contains(s, "world") {
        fmt.Println("s contains 'world'")
    } else {
        fmt.Println("s does not contain 'world'")
    }

	fmt.Printf("%s\n", strings.Replace(s, "l", "ooo", -1))
}