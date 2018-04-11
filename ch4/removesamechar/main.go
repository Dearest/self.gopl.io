package main

import (
	"fmt"
)

func removeSameChar(x []string) []string {
	var j int
	for i := 1; i < len(x); i++ {
		j = i - 1
		if x[i] == x[j] {
			x = remove(x, i)
		}
	}
	return x
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func main() {
	x := []string{"h", "h", "l"}
	x = removeSameChar(x)
	fmt.Printf("%q", x)
}

// out: ["h" "l"]
