package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
    stat := "File:`/usr/bin/ssh' Size: 318836 Blocks: 630 IO Block: 4096 Regular File Device: 302h/770d Inode: 178 Links: 1 Access: (0770/-rwxrwx---) Uid: ( 0/ root) Gid: ( 0/ root) Access: 2014-11-22 04:04:21.000000000 +0800 Modify: 2011-11-02 15:32:51.000000000 +0800 Change: 2020-09-26 04:33:47.000000000 +0800"
    file_info_en := regexp.MustCompile(`Size:\s(\d+).*?Access:\s\((\d+)/.*?Access:\s([\d\-\+\s:]+).*?Modify:\s([\d\-\+\s:]+).*?Change:\s([\d\-\+\s:]+)`)
    res := file_info_en.FindStringSubmatch(stat)
	fmt.Printf("%v\n", len(res))
    if len(res) > 1 {
        size, _ := strconv.Atoi(res[1])
        access := res[2]
		accessTime := res[3]
        modify := res[4]
        change := res[5]
        fmt.Println("Size:", size)
        fmt.Println("file Access:", access)
		fmt.Println("Access:", accessTime)
		var t time.Time
		var loc *time.Location
		var err error
		loc, err = time.LoadLocation("Local")
		if err != nil {
			fmt.Println("加载时区失败:", err)
			return
		}
		layout := "2006-01-02 15:04:05"
		t, _ = time.ParseInLocation(layout, accessTime, loc)
		fmt.Printf("%v\n", t.Unix())
		loc, err = time.LoadLocation("Local")
		t, err = time.ParseInLocation(layout, modify, loc)
        fmt.Println("Modify:", modify, t.Unix(), err)
		t, err = time.ParseInLocation(layout, change, loc)
		asciiValue := ""
		for i := 0; i < len(change); i++ {
			asciiValue += fmt.Sprintf("%d ", change[i])
		}
        fmt.Println("Change:", change, t.Unix(), err, asciiValue)
    }
}