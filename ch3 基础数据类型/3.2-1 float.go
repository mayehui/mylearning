package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216	// 1 << 24
	fmt.Println(f)		// 1.6777216e+07
	fmt.Println(f+1)	// 1.6777216e+07
	fmt.Println(f == f+1)		// "true"!
	// 当整数大于23bit能表达的范围时，float32表示出现误差

	// 浮点数的字面值可以直接写小数部分
	const e = 2.71828 // (approximately)

	const Avogadro = 6.02214129e23	// 阿伏伽德罗常数
	const Planck   = 6.62606957e-34	// 普朗克常数

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
