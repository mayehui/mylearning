// 2.3.2 指针

//x := 1
//p := &x         // p, of type *int, points to x
//fmt.Println(*p) // "1"
//*p = 2          // equivalent to x = 2
//fmt.Println(x)  // "2"

//var x, y int
//fmt.Println(&x == &x, &x == &y, &x == nil)
//				"true false false"

package main

import "fmt"

func main() {
	//x := 1
	//fmt.Println("x>>", x)
	//fmt.Println("&x>>", &x)
	//p := &x         // p, of type *int, points to x
	//fmt.Println("p>>", p)
	//fmt.Println("&p>>", &p)
	//fmt.Println("typeOf(*p)>>", reflect.TypeOf(*p))
	//fmt.Println(*p) // "1"
	//*p = 2          // equivalent to x = 2
	//fmt.Println(x)  // "2"

	var x, y int
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(&x == &x, &x == &y, &x == nil)

	//任何类型的指针的零值都是nil。
	var k *int
	fmt.Println(k)
}
