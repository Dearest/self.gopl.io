// 统计SHA256散列中不同的位数的总和
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n", s1, s2)
	fmt.Printf("the different count is %d", differentCount(s1, s2))
}

func differentCount(s1 [32]byte, s2 [32]byte) int {
	count := 0
	for i, v := range s1 {
		if s2[i] != v {
			count++
		}
	}
	return count
}
