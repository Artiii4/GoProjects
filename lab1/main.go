package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите первое число")
	var getIt string
	fmt.Scanf("%s\n", &getIt)
	firNum, err := strconv.ParseFloat(getIt, 64)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	fmt.Println("Выберите операцию(+, -, *,/)")
	var oper string
	fmt.Scanf("%s\n", &oper)
	if oper != "+" && oper != "*" && oper != "-" && oper != "/" {
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /")
		return
	}

	fmt.Println("Введите второе число")
	fmt.Scanf("%s\n", &getIt)
	secNum, err := strconv.ParseFloat(getIt, 64)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	result := firNum
	switch oper {
	case "+":
		result += secNum
	case "-":
		result -= secNum
	case "*":
		result *= secNum
	case "/":
		if secNum == 0 {
			fmt.Println("Ошибка: деление на 0 невозможно")
			return
		}
		result /= secNum
	}

	fmt.Println("Результат: ", firNum, oper, secNum, "=", result)
}
