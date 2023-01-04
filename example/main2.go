package main

import "fmt"

func main() {

	number := 5
	var p *int
	p = &number
	fmt.Println(p)
	*p = 10
	fmt.Println(p)
	fmt.Println(number)

	message := "i am learning golang"
	printMsg(&message) // & означает что мы передаём аргументом указатель (ссылку на ячейку в памяти)
	fmt.Println(message)
	fmt.Println(&message)
}

func printMsg(msg *string) { // * означает что мы принимаем не строку, а указатель(ссылку) на строку
	*msg += "\n it's very well"
	//fmt.Println(msg)
}
