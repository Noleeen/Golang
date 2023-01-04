package main

import (
	"errors"
	"fmt"
)

// функция init запускается при инициализации пакета перед main
//var msg string
//func init() {
//	msg = "what do you want?"
//}

func main() {
	f := anon()
	g := anon()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(anon())
	fmt.Println(prediction("wednesday"))
	fmt.Println(findMax(1, 153, 99, 1111))

	message, err := enterTheClub(15)
	if err != nil {
		//	//log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println(message)

}

func print(m string) {
	fmt.Println("функция принт пишет ", m)
}

func name(n string, age int) string {
	return fmt.Sprintf("прівет  %s ! Тебе %d лет", n, age)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Входи и развлекайся", nil
	} else if age >= 45 && age < 50 {
		return "тебе тут может быть скучно", nil
	} else if age == 50 || age > 50 {
		return "на баре забери приз", nil
	}
	return "ты ещё не дорос", errors.New("you are to young")
}

// конструкция switch case
func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "monday":
		return "you will find good job", nil
	case "tuesday":
		return "you will do it", nil
	case "wednesday":
		return "you will have great dinner", nil
	case "thursday":
		return "this day is fishDay", nil
	case "friday":
		return "ops...", nil
	default:
		return "недействительный день недели", errors.New("invalid day of the week")
	}
}

// итерация в цикле
func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}

// анонимная функция с замыканкием
func anon() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
