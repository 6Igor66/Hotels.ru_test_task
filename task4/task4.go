package main

import (
	"fmt"
)

func task4(target int) {
	fmt.Printf("\t")
	for i := 0; i < target; i++ {
		fmt.Printf("%d\t", i+1)
	}
	fmt.Printf("\n\n")
	for i := 0; i < target; i++ {
		fmt.Printf("%d\t", i+1)
		for j := 0; j < target; j++ {
			fmt.Printf("%d\t", (j+1)*(i+1))
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Println("Введите число: ")
	var target int
	fmt.Scan(&target)
	task4(target)
}
