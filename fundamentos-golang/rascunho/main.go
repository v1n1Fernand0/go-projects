package main

import "fmt"

func main() {
	// print a slice of data
	items := []int{10, 20, 40, 80}
	valor,erro := fmt.Println(items)
	fmt.Println(valor, erro)
}
