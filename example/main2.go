package main

import "fmt"

func main() {
	//var arr []string
	//arr[0] = "2"  будет ОШИБКА так как не инициализирована ячейка памяти для слайса

	arr := make([]string, 5, 15) // 15 - capasity  ЁМКОСТЬ !!!
	arr[2] = "7"

	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

}
