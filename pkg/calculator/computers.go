package calculate

func Computers(amount float64, duration int) float64 {

	switch duration {
	case 3:
		return (amount*(1*COMPUTERS_COMISSION))/100 + amount
	case 6:
		return (amount*(2*COMPUTERS_COMISSION))/100 + amount
	case 9:
		return (amount*(3*COMPUTERS_COMISSION))/100 + amount
	case 12:
		return (amount*(4*COMPUTERS_COMISSION))/100 + amount
	default:
		return 0
	}
}
