package main
 
import (
    "fmt"
    "log"
    "golang.org/x/crypto/bcrypt"
)
 
func main() {
    for {
       // 输入密码 获取 hash 值
        pwd := getPwd()
        hash := hashAndSalt(pwd)
       // 再次输入密码验证
        pwd2 := getPwd()
        pwdMatch := comparePasswords(hash, pwd2)
        fmt.Println("Passwords Match %v", pwdMatch)
    }
}
 
func getPwd() []byte {
    fmt.Println("Enter a password")
    var pwd string
    _, err := fmt.Scan(&pwd)
    if err != nil {
        log.Println(err)
    }
    return []byte(pwd)
}
 
func hashAndSalt(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}
 
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
    byteHash := []byte(hashedPwd)
 
	log.Println(hashedPwd, byteHash, plainPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}