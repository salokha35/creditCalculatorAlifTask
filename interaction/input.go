package Interaction

import (
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
		fmt.Println("Введите правильную сумму, сумма не может быть меньше или равно 0.")
		return
	}
	//So, before last step we get phone number and validate too
	fmt.Println("Введите номер телефона (бз кода страны и номер телефона) 992********* ")
	fmt.Scan(&phone)

	if len(phone) != 9 {
		fmt.Println("Неверный номер телефона")
		return
	}
	//The last step is asking duration of request
	fmt.Println("Введите срок кредита. Для вашего удобства вывели список допустимых месяцев для оформления кредита.")
	fmt.Scan(&duration)

	CategoryDistributions(category, amount, duration, phone)

}

func CategoryDistributions(category int, amount float64, duration int, phone string) {
	switch category {
	case 1:
		calculatedResult := Phones(amount, duration)
		println(int(calculatedResult))
	case 2:
		calculatedResult := Computers(amount, duration)
		println(int(calculatedResult))
	case 3:
		calculatedResult := Tvs(amount, duration)
		println(int(calculatedResult))
	}
}

func Computers(amount float64, duration int) float64 {

	switch duration {
	case 3:
		return (amount*(1*4))/100 + amount
	case 6:
		return (amount*(2*4))/100 + amount
	case 9:
		return (amount*(3*4))/100 + amount
	case 12:
		return (amount*(4*4))/100 + amount
	default:
		return 0
	}
}
func Phones(amount float64, duration int) float64 {
	switch duration {
	case 3:
		return (amount*(1*3))/100 + amount
	case 6:
		return (amount*(1*3))/100 + amount
	case 9:
		return (amount*(1*3))/100 + amount
	default:
		return 0
	}

}

func Tvs(amount float64, duration int) float64 {
	switch duration {
	case 3:
		return (amount*(1*5))/100 + amount
	case 6:
		return (amount*(2*5))/100 + amount
	case 9:
		return (amount*(3*5))/100 + amount
	case 12:
		return (amount*(4*5))/100 + amount
	case 18:
		return (amount*(5*5))/100 + amount
	default:
		return 0
	}
}

// func getInput() (string, error) {
// 	fmt.Print("Введите номер категории: ")

// 	userInput, err := reader.ReadString('\n')
// 	if err != nil {
// 		return "", err
// 	}

// 	userInput = strings.Replace(userInput, "\n", "", -1)

// 	return userInput, nil
// }
