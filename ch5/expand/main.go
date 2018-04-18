package main

import (
	"fmt"
	"strings"
)

func main() {
	f := gsub
	fmt.Println(expand("test$footest$foo", f))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func gsub(s string) string {
	return strings.ToUpper(s)
}
