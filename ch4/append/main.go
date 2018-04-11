// slice的动态扩容
package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	fmt.Printf("x cap=%d len=%d \t%v\n", cap(x), len(x), x)
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d len=%d \t%v\n", i, cap(y), len(y), y)
		x = y
	}
}
