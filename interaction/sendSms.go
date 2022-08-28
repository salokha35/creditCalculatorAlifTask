package Interaction

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessageToPhoneNumber(calculatedResult float64, amount float64, duration int, phone string) {

	message := "Детали платежа" +
		"\nЦена товара: " + fmt.Sprintf("%.1f\n", amount) +
		"Срок рассрочки: " + fmt.Sprintf("%d", duration) +
		"\nНомер телефона получателя: " + phone + "\nОбщяя сумма: " +
		fmt.Sprintf("%.1f\n", calculatedResult)

	accountSid := "ACffe25dc3fc93f2f818a96ed81d23205e"
	authToken := "f67fc049ea2f98ad86b9a36b84223193"
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken})
	params := &openapi.CreateMessageParams{}
	params.SetTo("+992" + phone)
	params.SetFrom("+13023356531")
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
