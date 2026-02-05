package main

import "fmt"

func main() {
	convertFrom, quantity, convertTo := userValueCurrency()
	convertMoney(convertFrom, quantity, convertTo)
}

func userValueCurrency() (float64, string, string) {
	const rub, usd, eur = "rub", "usd", "eur"
	var quantity float64
	var convertFrom, convertTo string
	for {
		fmt.Printf("Введите из какой валюты конвертировать (%s, %s, %s): ", rub, usd, eur)
		fmt.Scan(&convertFrom)
		if convertFrom == rub || convertFrom == usd || convertFrom == eur {
			break
		} else {
			fmt.Println("Вы ввели несуществующую валюту")
		}
	}
	for {
		fmt.Print("Введите количество валюты: ")
		_, err := fmt.Scan(&quantity)
		if err != nil || quantity < 1 {
			fmt.Println("Вы ввели, некоректную сумму")
		} else {
			break
		}
	}
	for {
		switch convertFrom {
		case rub:
			fmt.Printf("Введите в какую валюту конвертировать (%s, %s): ", usd, eur)
			_, err := fmt.Scan(&convertTo)
			if err != nil || (convertTo != usd && convertTo != eur) {
				fmt.Println("Вы ввели неверную валюту")
				continue
			}
		case usd:
			fmt.Printf("Введите в какую валюту конвертировать (%s, %s): ", rub, eur)
			_, err := fmt.Scan(&convertTo)
			if err != nil || (convertTo != rub && convertTo != eur) {
				fmt.Println("Вы ввели неверную валюту")
				continue
			}
		case eur:
			fmt.Printf("Введите в какую валюту конвертировать (%s, %s): ", rub, usd)
			_, err := fmt.Scan(&convertTo)
			if err != nil || (convertTo != rub && convertTo != usd) {
				fmt.Println("Вы ввели неверную валюту")
				continue
			}
		}
		break
	}
	return quantity, convertFrom, convertTo
}

func convertMoney(q float64, cf string, ct string) float64 {
	const usdToEur = 0.84
	const usdToRub = 77.02
	const eurToRub = usdToRub / usdToEur
	const rubToUsd = 1 / usdToRub
	const rubToEur = 1 / usdToEur
	const eurToUsd = 1 / usdToEur
	switch cf {
	case "rub":
		if ct == "usd" {
			result := q * rubToUsd
			fmt.Printf("При конвертации вы получите: %f usd", result)
			return result
		} else {
			result := q * rubToEur
			fmt.Printf("При конвертации вы получите: %f eur", result)
			return result
		}
	case "usd":
		if ct == "rub" {
			result := q * usdToRub
			fmt.Printf("При конвертации вы получите: %f rub", result)
			return result
		} else {
			result := q * usdToEur
			fmt.Printf("При конвертации вы получите: %f eur", result)
			return result
		}
	default:
		if ct == "usd" {
			result := q * eurToUsd
			fmt.Printf("При конвертации вы получите: %f usd", result)
			return result
		} else {
			result := q * eurToRub
			fmt.Printf("При конвертации вы получите: %f rub", result)
			return result
		}
	}
}
