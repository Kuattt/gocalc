package main

import (
	"fmt"
	// ... возможно, понадобится еще один пакет
)

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func divide(num1, num2 float64) float64 {
	return num1 / num2
}
func main() {

	// 2. Инициализируйте бесконечный цикл for { ... }
	for {
		// 3. Запрос чисел и оператора
		var num1 float64
		var num2 float64
		var operator string

		fmt.Print("Введите первое число: ")
		fmt.Scan(&num1)
		fmt.Print("Введите второе число: ")
		fmt.Scan(&num2)
		fmt.Print("Введите оператор (+, -, *, /): ")
		fmt.Scan(&operator)

		// 4. Реализуйте логику switch для выбора функции
		switch operator {
		case "+":
			fmt.Println(add(num1, num2))
			break
		case "-":
			fmt.Println(subtract(num1, num2))
			break
		case "*":
			fmt.Println(multiply(num1, num2))
			break
		case "/":
			if num2 == 0 {
				fmt.Println("Нельзя делить на ноль")
				break
			} else {
				fmt.Println(divide(num1, num2))
				break
			}
		default:
			fmt.Println("Неизвестный оператор")
		}
	}
}
