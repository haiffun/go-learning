package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 声明变量
var c, python, java bool

// 基础语法
func main() {
	fmt.Println("hello world")
	fmt.Printf("time: %s \n", time.Now().String())

	fmt.Println("random number: ", rand.Intn(10))

	fmt.Println(math.Pi)

	fmt.Println(add(1, 9))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(18))

	//var i, j int = 1, 2
	var i, j = 1, 2
	fmt.Println(i, j, c, python, java)

	i, j = 1, 2
	c, python, java := "c", true, 1
	fmt.Println(i, j, c, python, java)
}

func add(x int, y int) int {
	return x + y
}

// 支持多个返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 直接返回已命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
