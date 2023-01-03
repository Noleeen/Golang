package main

import (
	"errors"
	"fmt"
)

func main() {
	message, err := enterTheClub(15)
	if err != nil {
		//log.Fatal(err)
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
