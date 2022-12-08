package main

import "fmt"

func main() {

	fmt.Println("helo CODEMEN")
	var arr = [5]int{1: 3, 2: 5}
	fmt.Println("the value of the array", arr)

	//declaring a slice
	s := []int{1, 2, 3}
	var a int = cap(s)
	var b int = len(s)

	fmt.Println(a)
	fmt.Println(b)

	//slice_name := make([]type, length, capacity)
	s2 := make([]int, 5)
	fmt.Println("the value of the slice is ", s2)

	practice()
	example()
	multiple()

}
