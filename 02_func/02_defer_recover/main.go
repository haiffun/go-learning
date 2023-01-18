package main

import "fmt"

func main() {
	fmt.Println("before defer")

	// 延迟调用
	defer fmt.Println("defer1")
	fmt.Println("end defer")

	res, err := panicRecover()
	// 倒序延迟调用
	defer fmt.Println("defer2")
	fmt.Printf("res: %d \n", res)
	fmt.Printf("error: %s \n", err)
}

func panicRecover() (result int, err error) {
	defer func() {
		// recover, 捕获异常
		switch p := recover(); p {
		case nil:
			result = 1
		default:
			result = -1
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	panic("manual throw error")
	//return 0, nil
}
