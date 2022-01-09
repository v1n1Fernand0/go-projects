package main

import "fmt"

var n int

func main() {
	fmt.Println("O valor de n é", n)
}

func init() {
	n = 10
}
