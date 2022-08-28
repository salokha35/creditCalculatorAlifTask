package calculate

func Tvs(amount float64, duration int) float64 {
	switch duration {
	case 3:
		return (amount*(1*TVS_COMISSION))/100 + amount
	case 6:
		return (amount*(2*TVS_COMISSION))/100 + amount
	case 9:
		return (amount*(3*TVS_COMISSION))/100 + amount
	case 12:
		return (amount*(4*TVS_COMISSION))/100 + amount
	case 18:
		return (amount*(5*TVS_COMISSION))/100 + amount
	default:
		return 0
	}
}
