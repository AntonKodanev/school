package main

import "fmt"

func main() {
	const usdToEur = 0.84
	const usdToRub = 77.02
	myEur := 250.0
	eurToRub := myEur / 0.84 * 77.02
	fmt.Printf("%f евро в рублях: %f", myEur, eurToRub)
}
