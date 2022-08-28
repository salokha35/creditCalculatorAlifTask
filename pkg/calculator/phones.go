package calculate

func Phones(amount float64, duration int) float64 {
	switch duration {
	case 3:
		return (amount*(1*PHONES_COMISSION))/100 + amount
	case 6:
		return (amount*(1*PHONES_COMISSION))/100 + amount
	case 9:
		return (amount*(1*PHONES_COMISSION))/100 + amount
	default:
		return 0
	}

}
