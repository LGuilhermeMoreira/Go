package tax

import "time"

func CalculateTax(amount float64) float64 {
	if amount == 0 || amount < 0 {
		return 0
	}

	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20
	}

	return 5.0
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(3 * time.Second)

	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
