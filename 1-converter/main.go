package main

import "fmt"

func main() {
	const usdToEur = 0.84
	const usdToRub = 77.02
	const eurToRub = usdToRub / usdToEur
}

func userValueCurrency() (float64, string, string) {
	var quantity float64
	var convertFrom, convertTo string
	fmt.Print("Введите количество валюты: ")
	fmt.Scan(&quantity)
	fmt.Print("Введите из какой валюты конвертировать (rub, usd, eur): ")
	fmt.Scan(&convertFrom)
	fmt.Print("Введите в какую валюту конвертировать (rub, usd, eur): ")
	fmt.Scan(&convertTo)
	return quantity, convertFrom, convertTo
}

func convertMoney(q float64, cf string, ct string) (cm float64) {
	return 0.0
}
