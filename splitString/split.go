package splitString
import "strings"
 
//Newsplit 切割字符串
//example:
//abc,b=>[ac]
func Newsplit(str, sep string) []string {
    index := strings.Index(str, sep)
	var des []string
    for index > -1 {
        sectionBefor := str[:index]
        des = append(des, sectionBefor)
        str = str[index+1:]
        index = strings.Index(str, sep)
    }
    //最后1
    des = append(des, str)
    return des
}