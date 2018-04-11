// 一次遍历实现给定index的元素旋转 
package main

import "fmt"

func rotate(x []int, index int) []int {
	var res []int
	res = append(res, x[index+1:]...)
	res = append(res, x[index])
	res = append(res, x[:index]...)
	return res
}

func main() {
	a := []int{0, 1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("old: %v, expect: [3,4,2,1.0] got: %v\n", a, rotate(a, 2))
	fmt.Printf("old: %v, expect: [3,4,5,6,2,1], got: %v", b, rotate(b, 1))
}
