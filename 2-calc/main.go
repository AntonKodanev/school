package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	op := userOperationInput()
	str := userStringInput()
	sli := convertStringToSlice(str)
	calculate(op, sli)
}

func userOperationInput() string {
	var operation string
	for {
		fmt.Print("Введите арифметическую операцию (avg/sum/med): ")
		fmt.Scan(&operation)
		if operation == "sum" || operation == "avg" || operation == "med" {
			break
		} else {
			fmt.Println("Вы ввели неверную операцию")
		}
	}
	return operation
}

func userStringInput() string {
	fmt.Print("Введите любое количесто чисел разделенных запятой без пробелов (пример: 5,7,2,-2,-10): ")
	var stringInput string
	fmt.Scan(&stringInput)
	// fmt.Println(stringInput)
	return stringInput
}

func convertStringToSlice(str string) []int {
	parts := strings.Split(str, ",")
	result := make([]int, 0, len(parts))
	for _, s := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		result = append(result, n)
	}
	return result
}

func calculate(op string, sli []int) {
	switch {
	case op == "avg":
		fmt.Printf("Вы ввели %v\n", sli)
		sum := 0
		for _, value := range sli {
			sum += value
		}
		avg := float64(sum) / float64(len(sli))
		fmt.Printf("Среднее значение %v\n", avg)
	case op == "sum":
		fmt.Printf("Вы ввели %v\n", sli)
		sum := 0
		for _, value := range sli {
			sum += value
		}
		fmt.Printf("Сумма равна %v\n", sum)
	default:
		fmt.Printf("Вы ввели %v\n", sli)
		tmp := make([]int, len(sli))
		copy(tmp, sli)
		sort.Ints(tmp)
		n := len(tmp)
		if len(sli) == 0 {
			fmt.Printf("Медиана равно 0")
		} else if n%2 != 0 {
			fmt.Printf("Медиана равна %v", tmp[n/2])
		} else {
			fmt.Printf("Медиана равна %v", float64(tmp[n/2-1]+tmp[n/2])/2)
		}
	}
}
