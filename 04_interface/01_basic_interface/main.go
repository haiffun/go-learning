package main

import "fmt"

type F interface {
	Str()
}

type Fs struct {
}

// Str *Fs有F所有方法 -> *Fs 实现了接口F
func (*Fs) Str() {
	fmt.Println("=== fs.str ===")
}

func main() {
	var f1 F     // 类型和值都是nil
	f2 := new(F) // 类型是*main.f, 值是nil
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
	fmt.Println(f1 == nil)
	fmt.Println(f2 == nil)

	f1 = new(Fs)
	fmt.Printf("%T\n", f1)
	f1.Str()
}
