package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"github.com/PuerkitoBio/goquery"
)


//OptionFun 函数的参数
type OptionFun func(client *http.Client)

//with
func withTimeout(duration time.Duration) OptionFun {
	return func(client *http.Client) {
		client.Timeout = duration * time.Microsecond
	}
}

func NewByOption(opts ...OptionFun) http.Client {
	client := http.Client{}
	for _, opt := range opts {
		opt(&client)
	}
	return client
}

func main() {
	GetTest()
	//PostTest()
	//PutTest()
	//DeleteTest()
	//TimeoutTest()
}

func GetTest()  {
	reqest,err := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http.NewRequest :",err)
	}
	//创建http客户端
	var c = NewByOption()
	response,err := c.Do(reqest)
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("response.StatusCode :",response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	sBody := string(body)
	//bodyMap := make(map[string]interface{})
	//err = json.Unmarshal(body,&bodyMap)
	fmt.Println(sBody)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println("goquery.NewDocumentFromReader :",err)
	}
	doc.Find
}
func PostTest()  {
	//POST
	//在 POST 方式一般常用的为 2 中，
	//
	//通过 kv 形式传送，例如 form-data 和 x-www-form-urlencoded
	//通过 json 形式传送，例如 application/json

	//kv 形式
	//payload := strings.NewReader("a=111")
	//json 形式
	payload := strings.NewReader("{\"name\":\"张三\"}")
	reqest,err := http.NewRequest(http.MethodPost,"http://127.0.0.1:8080/user",payload)
	if err != nil {
		fmt.Println("http.NewRequest",err)
	}
	reqest.Header.Add("Content-Type", "application/json")
	//创建http客户端
	var c = NewByOption()
	response,err := c.Do(reqest)
	if err != nil {
		fmt.Println("c.Do",err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("response.StatusCode :",response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body,&bodyMap)
	fmt.Println(bodyMap)
}
func PutTest()  {
	//由于 net/http 没有提供简化的 PUT 请求，这里需要使用 http.NewRequest 来创建请求
	payload := strings.NewReader("{\"name\":\"张三\"}")
	reqest,err := http.NewRequest(http.MethodPut,"http://127.0.0.1:8080/user",payload)
	if err != nil {
		fmt.Println("http.NewRequest",err)
	}
	reqest.Header.Add("Content-Type", "application/json")

	//创建http客户端
	var c = NewByOption()
	response,err := c.Do(reqest)
	if err != nil {
		fmt.Println("c.Do",err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("response.StatusCode :",response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body,&bodyMap)
	fmt.Println(bodyMap)
}
func DeleteTest()  {
	payload := strings.NewReader("{\"name\":\"张三\"}")
	reqest,err := http.NewRequest(http.MethodDelete,"http://127.0.0.1:8080/user",payload)
	if err != nil {
		fmt.Println("http.NewRequest",err)
	}
	reqest.Header.Add("Content-Type", "application/json")

	//创建http客户端
	var c = NewByOption()
	response,err := c.Do(reqest)
	if err != nil {
		fmt.Println("c.Do",err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("response.StatusCode :",response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body,&bodyMap)
	fmt.Println(bodyMap)
}
func TimeoutTest()  {
	reqest,err := http.NewRequest(http.MethodGet,"https://www.baidu.com",nil)
	if err != nil {
		fmt.Println("http.NewRequest :",err)
	}
	//创建http客户端
	//var c = NewByOption()	//不报错
	//var c = NewByOption(withTimeout(3))	//不报错
	var c = NewByOption(withTimeout(1))	//报错
	response,err := c.Do(reqest)
	if err != nil {
		fmt.Println("c.Do(reqest) err ======",err)
		return
	}
	//fmt.Println("response.Body==",response.Body)
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("response.StatusCode :",response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body,&bodyMap)
	fmt.Println(bodyMap)
}
