package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
    Big   = 1 << 100
    Small = Big >> 99
)

// 基本数据类型
func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)

    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)

    //var x int = 42
    //var y float64 = float64(x)
    //var u uint = uint(y)

    x := 42 // 类型推导
    y := float64(x)
    u := uint(y)
    fmt.Printf("%T %v %T %v %T %v\n", x, x, y, y, u, u)

    const Pi = 3.14 // 常量
    fmt.Println(Pi)

    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}

func needInt(x int) int {
    return x*10 + 1
}

func needFloat(x float64) float64 {
    return x * 0.1
}
