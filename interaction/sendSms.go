package Interaction

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

/*
Why did I use twilio?
At the beginning, I tried through osonsms, but they only provide services for individual entrepreneurs.
Then I tried Russian services, but they have a limited trial version.
According to the experience of some colleagues, I wanted to try the textmagic, but there were problems with registration.
Then I found twilio, and I think it fits best.
But there is one difficulty, before sending SMS, you need to add this number to your contact list.
If necessary, write to me in a telegram: solievsalah
*/

func SendMessageToPhoneNumber(calculatedResult float64, amount float64, duration int, phone string) {
	//Next, we get the variables and compose the text of our SMS message
	message := "Детали платежа:" +
		"\nЦена товара: " + fmt.Sprintf("%.1f\n", amount) +
		"Срок рассрочки: " + fmt.Sprintf("%d", duration) +
		"\nНомер телефона получателя: " + phone + "\nОбщяя сумма: " +
		fmt.Sprintf("%.1f\n", calculatedResult)

	//Next, we enter the data for authorization in the Twilio service for sending SMS messages.
	accountSid := "ACffe25dc3fc93f2f818a96ed81d23205e"
	authToken := "f67fc049ea2f98ad86b9a36b84223193"
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken})
	params := &openapi.CreateMessageParams{}
	params.SetTo("+992" + phone)
	params.SetFrom("+13023356531")
	params.SetBody(message)

	//set handler
	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
