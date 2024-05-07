package main

import "fmt"

func task1(value int) string {
	ost := value % 10
	ostDec := value % 100

	if ost == 1 && ostDec != 11 {
		return fmt.Sprintf("%d компьютер", value)
	}

	if ost >= 2 && ost <= 4 && (ostDec < 10 || ostDec >= 20) {
		return fmt.Sprintf("%d компьютера", value)
	}

	return fmt.Sprintf("%d компьютеров", value)
}

func main() {
	var val int
	fmt.Println("Введите число: ")
	fmt.Scan(&val)
	str := task1(val)
	fmt.Println(str)
}
