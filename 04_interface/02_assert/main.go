package main

import "fmt"

type F1 interface {
	m1()
}

type F2 interface {
	F1
	m2()
}

type S1 struct {
	v int
}

func (s1 *S1) m1() {
	fmt.Println("S1.M1")
}

type S2 struct {
}

func (s2 *S2) m1() {
	fmt.Println("S2.M1")
}

func (s2 *S2) m2() {
	fmt.Println("S2.M2")
}

func main() {
	var f1 F1
	f1 = new(S2)

	// 类型是F1，动态类型是*S2
	f1.m1()

	r1 := f1.(*S2)
	// 类型是*S2
	r1.m1()
	r1.m2()
	fmt.Printf("%T\n", r1)
	fmt.Println(r1)

	r2 := f1.(F2)

	// 类型是F2，动态类型是*S2
	r2.m1()
	r2.m2()
	fmt.Printf("%T\n", r2)
	fmt.Println(r2)
}
