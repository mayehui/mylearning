// 2.4 赋值
package main

import (
	"fmt"
)

func main() {
	def240()
	def241()
	fmt.Println("hello word")
}

func def240() {
	x = 1		//命名变量的赋值
	*p = true	//通过指针间接赋值
	person.name = "bob"	//结构体字段赋值
	count[x] = count[x] * scale //数组、slice或map的元素赋值
	// count[x] *= scale
}

func def241() {
	x, y = y, x
	a[i], a[j] = a[j], a[i]
	gcd(1,5)
	fib(8)
	//元组赋值
	//i, j, k = 2, 3, 5
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y. x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}