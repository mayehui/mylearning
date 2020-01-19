// gopl.io/ch1/echo1
// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	var s, s1, sep string
	s1 = os.Args[0]
	for i := 0; i < len(s1) ;i++  {
		s += sep + string(s1[i])
		sep = ""
	}
	fmt.Println(s)
}