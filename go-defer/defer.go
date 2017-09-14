package go_defer

import "fmt"

func deferRet(x, y int) (z int) {
	defer fmt.Println("defer1 ...")
	defer fmt.Println("defer2 ...")

	z = x + y
	fmt.Println("in deferRet()")
	return z
}

func Run() {
	i := deferRet(1, 1)
	fmt.Println(i)
}