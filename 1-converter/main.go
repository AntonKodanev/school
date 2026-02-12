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

func convertMoney(q float64, cf string, ct string) {
	const usdToEur = 0.84
	const usdToRub = 77.02
	const eurToRub = usdToRub / usdToEur
	const rubToUsd = 1 / usdToRub
	const rubToEur = 1 / eurToRub
	const eurToUsd = 1 / usdToEur

	type Money = map[string]map[string]float64
	userMoney := Money{"usd": {"eur": usdToEur, "rub": usdToRub}, "eur": {"rub": eurToRub, "usd": eurToUsd}, "rub": {"usd": rubToUsd, "eur": rubToEur}}

	for k, v := range userMoney {
		if k == cf {
			for u, p := range v {
				if ct == u {
					result := q * p
					fmt.Printf("При конвертации вы получите: %f %v", result, u)
				}
			}
		}
	}
}
