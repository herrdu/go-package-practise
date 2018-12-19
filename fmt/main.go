package main

import "fmt"

type Website struct {
	Name string
}

// 定义结构体变量
var site = Website{Name: "studygolang"}

func main() {
	fmt.Printf("%T", site)
}
