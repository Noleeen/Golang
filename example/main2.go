package main

import "fmt"

// МАПЫ-----------------------------------------------------
func main() {
	users := map[string]int{
		"Nina":  17,
		"Sasha": 15,
		"Petya": 21,
		"Lesya": 18,
		"Aryna": 10,
	}

	age, exists := users["Sasha"]

	if exists {
		fmt.Println(exists, age, "\n")
	} else {
		fmt.Println("this name isn't in map\n")
	}

	users["Kate"] = 27     // добавляем новый ключ и значение в мапу
	delete(users, "Petya") // удаление записи по ключу
	// ИТЕРАЦИЯ ПО МАПАМ
	for key, value := range users {
		fmt.Println(key, value)
	}
}

//----------------------------------------------------------

// итерация масивов и слайсов-------------------------------
//func main() {
//
//	defer hendlerPanic()
//
//	array := []string{
//		"m1",
//		"m2",
//		"m3",
//		"m4",
//		"m5",
//	}
//
//	for i, value := range array {
//		fmt.Println("index ", i)
//		fmt.Println("value ", value)
//	}
//
//	array[5] = "sd"
//}
//
//func hendlerPanic() {
//	if r := recover(); r != nil {
//		fmt.Println(r)
//	}
//	fmt.Println("handlerPanic was run!")
//}
//-------------------------------------------------------

// MATRIX  двумерный массив -----------------------------
//func main() {
//	matrix := make([][]int, 10)
//
//	for x := 0; x < 10; x++ {
//		matrix[x] = make([]int, 10)
//		for y := 0; y < 10; y++ {
//			matrix[x][y] = x
//		}
//		fmt.Println(matrix[x])
//	}
//}
//---------------------------------------

// УКАЗАТЕЛИ ----------------------------------------

//func main() {
//
//	number := 5
//	var p *int
//	p = &number
//	fmt.Println(p)
//	*p = 10
//	fmt.Println(p)
//	fmt.Println(number)
//
//	message := "i am learning golang"
//	printMsg(&message) // & означает что мы передаём аргументом указатель (ссылку на ячейку в памяти)
//	fmt.Println(message)
//	fmt.Println(&message)
//}
//
//func printMsg(msg *string) { // * означает что мы принимаем не строку, а указатель(ссылку) на строку
//	*msg += "\n it's very well"
//	//fmt.Println(msg)
//}
//--------------------------------------------------

//// СЛАЙСЫ ------------------------------
//func main() {
//	//var arr []string
//	//arr[0] = "2"  будет ОШИБКА так как не инициализирована ячейка памяти для слайса
//
//	arr := make([]string, 5, 15) // 15 - capasity  ЁМКОСТЬ !!!
//	arr[2] = "7"
//
//	fmt.Println(arr)
//	fmt.Println(len(arr))
//	fmt.Println(cap(arr))
//
//}
//------------------------------------------------
