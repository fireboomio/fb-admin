package main

import (
	"fmt"
	"simple-casdoor/object"
)

func main() {
	// 初始化 xorm 映射
	if err := object.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	object.InitAdapter()

	r := NewRouter()

	r.Logger.Fatal(r.Start(":9825"))
}
