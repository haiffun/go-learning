// 函数
package main

import (
	"fmt"
	"go-learning/02_func/utils"
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

	subNum = utils.Sum(num, 2)
	fmt.Println(subNum)
}

func sub(val1 int, val2 int) (sum int) {
	sum = val1 + val2
	return
}

func initNum() int {
	fmt.Println("global parameter...")
	return 0
}
