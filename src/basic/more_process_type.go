package main

import (
    "fmt"
)

// 结构体: 即一组字段
type Vertex struct {
    X int // 结构体字段
    Y int
}

var (
    v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
    v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
    v3 = Vertex{}      // X:0 Y:0
    p1 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

// 指针、结构体、数组(切片)
func main() {
    i, j := 42, 2701

    // 指针
    p := &i         // 指向 i
    fmt.Println(*p) // 通过指针读取 i 的值
    *p = 21         // 通过指针设置 i 的值
    fmt.Println(i)  // 查看 i 的值

    p = &j         // 指向 j
    *p = *p / 37   // 通过指针对 j 进行除法运算
    fmt.Println(j) // 查看 j 的值

    // 结构体
    v := Vertex{1, 2}
    fmt.Println(v)
    fmt.Println(v.X)

    n := &v
    n.X = 1e9 // 隐式间接引用
    fmt.Println(v)

    fmt.Println(v1, p1, v2, v3)

    // 数组
    fmt.Println("array ====================")

    var a [2]string
    a[0] = "hello"
    a[1] = "array"
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    // 切片
    fmt.Println("array split ====================")

    s := primes[1:4] // [1,4)
    fmt.Println(s)

    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

    nameA := names[0:2] // 切片不存储数据, 只是描述数组的一段信息
    nameB := names[1:3]
    fmt.Println(nameA, nameB)

    nameB[0] = "XXX" // 更改切片元素会修改底层数组中的对应元素
    fmt.Println(nameA, nameB)
    fmt.Println(names)

    // 切片下界默认为0, 上界默认为数组长度
    fmt.Println("array1 ====================")

    array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

    // 以下[0:10] [:10] [0:] [:]行为等价
    arraySplit := array[0:10]
    fmt.Println(arraySplit)

    arraySplit = array[:10]
    fmt.Println(arraySplit)

    arraySplit = array[0:]
    fmt.Println(arraySplit)

    arraySplit = array[:]
    fmt.Println(arraySplit)

    fmt.Println("array2 ====================")
    array2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    // 截取切片使其长度为 0
    array2 = array2[:0]
    printSlice(array2)
    // 拓展其长度
    array2 = array2[:4]
    printSlice(array2)
    // 舍弃前两个值
    array2 = array2[2:]
    printSlice(array2)

    fmt.Println("array3 nil ====================")
    var array3 []int // 切片的零值是 nil, 长度和容量为0且没有底层数组
    printSlice(array2)
    if array3 == nil {
        fmt.Println("array3 is nil!")
    }

    // make函数, 创建动态数组
    fmt.Println("array4 make ====================")

    array4a := make([]int, 0, 5) // 分配一个元素为零值的数组并返回一个引用它的切片, 长度为0, 容量为5
    printSlice(array4a)

    array4b := make([]int, 5) // 等价于make([]int, 5, 5)
    printSlice(array4b)

    array4c := array4b[:2]
    printSlice(array4c)

    array4d := array4c[2:5]
    printSlice(array4d)

    // 向切片追加元素 append(s []T, vs ...T) []T
    fmt.Println("array5 append ====================")

    var array5 []int
    printSlice(array5)

    // 添加一个空切片
    array5 = append(array5, 0)
    printSlice(array5)

    // 这个切片会按需增长
    array5 = append(array5, 1)
    printSlice(array5)

    // 可以一次性添加多个元素
    array5 = append(array5, 2, 3, 4)
    printSlice(array5)

    // for range遍历切片或映射, 每次迭代返回当前元素下标和对应元素的副本
    fmt.Println("range ====================")

    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    pow = make([]int, 10)
    for i := range pow { // 第二个变量可忽略
        pow[i] = 1 << uint(i) // == 2**i
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }

    // 映射map
    fmt.Println("map ====================")

    var m map[string]Vertex
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40, -74,
    }
    fmt.Println(m["Bell Labs"])

    m = map[string]Vertex{
        "Bell Labs": {
            40, -74,
        },
        "Google": {
            37, -122,
        },
    }
    fmt.Println(m)

    // 修改映射
    fmt.Println("map modify ====================")

    m2 := make(map[string]int)

    m2["Answer"] = 42 // 插入或修改元素
    fmt.Println("The value:", m2["Answer"])

    m2["Answer"] = 48
    fmt.Println("The value:", m2["Answer"])

    delete(m2, "Answer") // 删除元素
    fmt.Println("The value:", m2["Answer"])

    v2, ok := m2["Answer"] // 双赋值检测键是否存在
    fmt.Println("The value:", v2, "Present?", ok)

    // 函数值
    fmt.Println("function value ====================")

    makeMap := func() map[string]int {
        fmt.Println("make map")
        return make(map[string]int)
    }
    fmt.Println(makeMap())

    // 函数闭包, 闭包是一个函数值，它引用了其函数体之外的变量
    fmt.Println("function closure ====================")
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x) // len取数组长度, cap取数组容量
}

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
