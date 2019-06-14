# Go字符串加密实现

## 场景

Go字符串的加密场景有许多这里是最常见的场景：

- 用户注册：用户输入的密码信息需要被加密；
- 用户登录：用户登录需要和数据库中的密码进行确认对比;
- 重要信息的保密:  其实任何重要的信息都需要加密。

当然，字符串加密算法有很多，这里之所以采用bcrypt是模仿laravel中默认的加密算法，并且golang提供了原生的加密以及对比算法，非常安全，特别容易编码实现。

## 使用的包

> golang.org/x/crypto/bcrypt

> 安装: go get golang.org/x/crypto/bcrypt

> 文档所在地址: https://godoc.org/golang.org/x/crypto/bcrypt

包内预定义常量

```go
const (
    MinCost     int = 4  // 最小密文加密长度
    MaxCost     int = 31 // 最大密文加密长度
    DefaultCost int = 10 // 默认密文加密长度
)
```

func GenerateFromPassword: 创建加密密码

```go
func GenerateFromPassword(密码明文 []byte, 密文加密长度 int) ([]byte, error)
```

func CompareHashAndPassword:对比明文和密码

```go
func CompareHashAndPassword(密文, 明文 []byte) error
```

## 完整代码实例

```go
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

```
