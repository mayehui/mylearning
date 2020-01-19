// gopl.io/ch1/echo2
// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	var s1 = os.Args[0]
	fmt.Println(s1)
	s, sep := "",""
	for _, arg := range s1 {
		fmt.Println(reflect.TypeOf(arg))
		//s += sep + arg
		sep += " "
	}
	fmt.Println(s)
}