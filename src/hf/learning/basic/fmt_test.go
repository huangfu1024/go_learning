package basic

import "fmt"

func fmtTest() {
	var name, age string
	fmt.Printf("输入名字, \n")
	fmt.Scanln(&name)
	fmt.Printf("输入年龄, \n")
	fmt.Scanln(&age)
	fmt.Printf("你的名字是: %a, %b", name, age)
}
