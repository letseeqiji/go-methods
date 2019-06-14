package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
	password := "test123456"

	// 加密密码: 但是请注意，返回的并不是string类型，需要转换
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
    		fmt.Println(err)
	}

	// 这里是要入库的密码
	encodePW := string(hash)
	fmt.Println(encodePW)

	// 登录时验证密码的方法
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(password))
	if err != nil {
		fmt.Println("密码错误")
	} else {
		fmt.Println("密码正确")
	}
}
