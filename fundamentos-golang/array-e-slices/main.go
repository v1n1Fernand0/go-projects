package main

import (
	"fmt"
	"reflect"
)

func main() {
	//array
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"1", "teste", "3", "4", "5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	//slice
	slice := []int{10}
	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(array3)) // comparando os tipos

	slice = append(slice, 28) // Atribuição de valores ao slice
	fmt.Println(slice)

	slice2 := array2[1:3] // criando um slice apartir de um array
	fmt.Println(slice2)

	//arrays Internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
