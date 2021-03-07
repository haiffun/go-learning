package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

// 流程控制: for/if/else/switch/defer
func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += 1
    }
    fmt.Println("for loop:", sum)

    for sum < 100 {
        sum++
    }
    fmt.Println("for loop2:", sum)

    for sum < 200 {
        sum++
    }
    fmt.Println("for loop3:", sum)

    // 无限循环
    //for {
    //	fmt.Println("for loop4:", sum)
    //}

    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )

    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Printf("%s.\n", os)
    }

    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }

    defer fmt.Println("world") //defer 语句会将函数推迟到外层函数返回之后执行
    fmt.Println("hello")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i) //defer 函数会被压到栈中, 以后进先出形式调用
    }
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}
