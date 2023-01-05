package main

import "fmt"

// КОНСТРУКТОР---------------------------------------------
// то же самое как и в нижней структуре:
func newUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height,
	}
}
func main() {
	user1 := newUser("Vasya", "Male", 72, 82, 176)
}

// ЕЩЁ ОДИН ВИД ОБЪЯВЛЕНИЯ КОНСТРУКТОРА
type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

//-----------------------------------------------------------------------

// СТРУКТУРЫ -----------------------------------------------
//	user := struct {
//		name   string
//		age    int
//		sex    string
//		weight int
//		height int
//	}{"Vasya", 72, "Male", 82, 176}   // такой вариант объявления структуры нельзя переиспользовать, поэтому используют чаще другой (СМ. НИЖЕ)

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func main() {
	user1 := User{"Vasya", 72, "Male", 82, 176}
	user2 := User{"Nina", 65, "Female", 90, 173}

	fmt.Println("\n", user1, "\n")
	fmt.Println(user1.name, user1.sex, "\n")
	fmt.Printf("%+v\n", user1)
	fmt.Printf("%+v\n", user2)

}

// ---------------------------------------------------------

// МАПЫ-----------------------------------------------------
//func main() {
//	users := map[string]int{
//		"Nina":  17,
//		"Sasha": 15,
//		"Petya": 21,
//		"Lesya": 18,
//		"Aryna": 10,
//	}
//
//	age, exists := users["Sasha"]
//
//	if exists {
//		fmt.Println(exists, age, "\n")
//	} else {
//		fmt.Println("this name isn't in map\n")
//	}
//
//	users["Kate"] = 27     // добавляем новый ключ и значение в мапу
//	delete(users, "Petya") // удаление записи по ключу
//	// ИТЕРАЦИЯ ПО МАПАМ
//	for key, value := range users {
//		fmt.Println(key, value)
//	}
//}

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
