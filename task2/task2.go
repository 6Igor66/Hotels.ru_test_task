package main

import (
	"fmt"
)

func task2(arr []int) []int {
	ans := []int{}
	mp := make(map[int]int)
	for i := range arr {
		for j := 2; j <= arr[i]; j++ {
			if arr[i]%j == 0 {
				mp[j]++
			}
		}
	}

	for k, v := range mp {
		if v == len(arr) {
			ans = append(ans, k)
		}
	}

	return ans

}

func main() {
	fmt.Println("Введите длину массива: ")
	var k int
	fmt.Scan(&k)
	fmt.Println("Введите элементы массива: ")
	arr := []int{}
	for i := 0; i < k; i++ {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	ans := task2(arr)
	fmt.Printf("Общие делители для указанных чисел: \n%v", ans)
}
