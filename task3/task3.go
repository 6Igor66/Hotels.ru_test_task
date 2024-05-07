package main

import "fmt"

func task3(low, high int) *[]int {
	arr := []int{}
	ans := []int{}
	for i := range high + 1 {
		arr = append(arr, i)
	}
	arr[1] = 0
	i := 2
	for i <= high {
		if arr[i] != 0 {
			if i >= low {
				ans = append(ans, i)
			}
			j := i + i
			for j <= high {
				arr[j] = 0
				j = j + i
			}
		}
		i += 1
	}
	return &ans

}

func main() {
	var low, high int
	fmt.Println("Введите диапазон значений (минимальное и максимальное): ")
	fmt.Scan(&low, &high)
	ans := task3(low, high)
	fmt.Println(*ans)
}
