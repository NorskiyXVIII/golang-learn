/*
\	КОНТРОЛЬНАЯ РАБОТА
\   Получаем от пользователя разный ввод
\   Если 0b, то бинари в десятичные, если 0, то в восьмеричные и т.д
*/

package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Print("Введите [НЕ ДЕСЯТИЧНОЕ] число или введите 'stop' для выхода: ")

	var input string

	for {

		fmt.Scanln(&input)
		input = strings.ToLower(input)

		if input == "stop" || input == "стоп" {
			break
		}

		num := new(big.Int)
		procedNum, razr := procNum(input)

		switch razr {
		case 1:
			num.SetString(procedNum, 16)
		case 2:
			num.SetString(procedNum, 2)
		case 3:
			num.SetString(procedNum, 8)
		case 4:
			fmt.Println("Warn: Вы ввели десятичное значение")
			num.SetString(procedNum, 10)
		default:
			fmt.Println("Фейл")
			continue
		}

		fmt.Println("Итоговое значение: ", num)
	}
}

func procNum(numStr string) (string, int) {
	if numStr[0] == '0' && numStr[1] == 'x' {
		return strings.TrimPrefix(numStr, "0x"), 1
	} else if numStr[0] == '0' && numStr[1] == 'b' {
		return strings.TrimPrefix(numStr, "0b"), 2
	} else if numStr[0] == '0' {
		return strings.TrimPrefix(numStr, "0"), 3
	} else {
		return numStr, 4
	}
}
