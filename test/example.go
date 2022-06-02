package main

import (
	"fmt"
	"github.com/PeterYangs/consistent"
)

func main() {

	h := consistent.New()

	//添加节点
	h.Add("192.168.1.10")
	h.Add("192.168.1.11")
	h.Add("192.168.1.12")
	h.Add("192.168.1.13")
	h.Add("192.168.1.14")
	h.Add("192.168.1.15")
	h.Add("192.168.1.16")

	//hash取模
	fmt.Println(h.Get("https://www.baidu.com"))
}
