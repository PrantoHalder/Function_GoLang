package main

import "fmt"

func multiple() {

	day := 5

	switch day {

	case 1, 2, 5:
		fmt.Println("pranto")
	case 3, 4:
		fmt.Println("halder")
	default:
		fmt.Println("pranto hadler")

	}
}
