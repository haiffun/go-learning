// 函数
package main

import (
	"fmt"
)

// 执行顺序：全局参数初始化 -> init -> main
var num = initNum()

// init函数先执行
func init() {
	fmt.Println("init func...")
}

func main() {
	fmt.Println("main func...")

	// 函数也是一种数据类型
	subFunc := sub
	subNum := subFunc(num, 2)
	fmt.Println(subNum)

	subFunc2 := sub2
	// 函数类型/函数值
	fmt.Printf("%T\n", subFunc)
	fmt.Printf("%T\n", subFunc2)

	// 函数值的零值是nil，函数值只能与nil比较，函数值之间不能比较
	if subFunc != nil {
		subFunc(1, 2)
	}

	subNum = Sum(num, 2)
	fmt.Println(subNum)
}

func sub(val1 int, val2 int) (sum int) {
	sum = val1 + val2
	return
}

func sub2(val1 int, val2 int) (sum int) {
	sum = val1 + val2
	return
}

func initNum() int {
	fmt.Println("global parameter...")
	return 0
}
