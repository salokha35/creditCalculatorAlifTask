package Interaction

import (
	calculate "creditCalculatorAlifTask/pkg/calculator"
	"fmt"
)

func InputData() {

	var (
		category int
		amount   float64
		duration int
		phone    string
	)
	//In this part of code, we ask from user category of his request
	fmt.Println("Выберите категорию товара:\n" + "1. Смартфон\n" + "2. Компьютер\n" + "3. Телевизор")
	fmt.Scan(&category)

	//Then we ask amount of request and validate it, to get correct amount
	fmt.Println("Введите сумму продукта")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Введите правильную сумму,сумма должна быть больше 0.")
		return
	}

	//So,before last step is asking duration of request
	fmt.Println("Введите срок кредита.")
	fmt.Scan(&duration)

	//The last step we get phone number and validate too
	fmt.Println("Введите номер телефона (без кода страны и номер телефона) 992********* ")
	fmt.Scan(&phone)

	if len(phone) != 9 {
		fmt.Println("Неверный номер телефона")
		return
	}

	CategoryDistributions(category, amount, duration, phone)

}

func CategoryDistributions(category int, amount float64, duration int, phone string) {
	switch category {
	case 1:
		calculatedResult := calculate.Phones(amount, duration)
		SendMessageToPhoneNumber(calculatedResult, amount, duration, phone)
	case 2:
		calculatedResult := calculate.Computers(amount, duration)
		SendMessageToPhoneNumber(calculatedResult, amount, duration, phone)
	case 3:
		calculatedResult := calculate.Tvs(amount, duration)
		SendMessageToPhoneNumber(calculatedResult, amount, duration, phone)
	}
}
