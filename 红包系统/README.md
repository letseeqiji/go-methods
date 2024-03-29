[![Language](https://img.shields.io/badge/Language-Golang-blue.svg)](https://github.com/letseeqiji/git-helper)
[![Build Status](https://travis-ci.org/bilibili/kratos.svg?branch=master)](https://github.com/letseeqiji/git-helper)

# 红包系统关键点

这里讨论红包系统的关键算法。  

> 抢红包到很多人的喜欢，但是红包的算法和系统设计如何，还需要认真讨论。

## 目标

> 致力于提供更加方便快捷的操作方式，节省更多的时间去创造更具价值的东西。

### **部分代码**

```go
......
#二倍均值算法
// 一分钱
var min int64 = 1
func DoubleAverge(count, amount int64) int64 {
    // 如果总数为1 直接返回
    if count == 1 {
        return amount
    }
    // 计算红包最大可用值：红包的剩余金额 - 一分钱 * 红包总数
    max := amount - min * count
    // 计算最大可用平均值
    avg := max / count
    // 防止出现0
    avg = 2 * avg + min
    // 生成随机红包
    rand.Seed(time.Now().UnixNano())
    x := rand.Int63n(avg) + min
    return x
}
......
```

## 快速开始

### 获取

```shell
git clone https://github.com/letseeqiji/oneinstall.git
cd oneinstall/红包
```

------

## 文档

[简体中文](https://github.com/letseeqiji/oneinstall/blob/master/golang/README.md)

------

*Please report bugs, concerns, suggestions by issues, or join QQ 962310113to discuss problems around source code.*
