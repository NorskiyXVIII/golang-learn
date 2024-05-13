package main

import (
	"fmt"
	"math/big" // большие числа
	"strings"  // для функции ToLower
)

func main() {
	fmt.Print("Enter hex number or 'stop' to exit: ")

	var input string

	// Если в for нет условия, то автоматически for будет бесконечным
	for {
		// Scanln в пространстве имен fmt. В отличии от scanf, вернет в любом случае строку, ибо ввод от пользователя - строка
		// если видит символ пробела, то тогда берет первое слово и заносит в переменную, все остальное вычищает
		fmt.Scanln(&input)

		//fmt.Print("Ввод пользователя: ", user_string, "\n")

		// делаем все буквы строчными
		input = strings.ToLower(input)

		if input == "stop" {
			break // выход из цикла
		}

		// функция new размечает память в ОЗУ(захватывает в памяти свободное место)
		i := new(big.Int)

		// метод SetString берет строку и конвертирует число по основанию base. args - (string, base)
		i.SetString(processHex(input), 16)

		// _ пропускает возвращенное значение
		/*_, ok := i.SetString(processHex(input), 16)

		if !ok {
			fmt.Println("FAIL")
		} else {
			fmt.Println(i)
		}*/

		// можно сделать инициализацию в if как в for
		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("FAIL")
			continue
		}

		fmt.Println(i)
	}
}

func processHex(hexStr string) string {
	// 0xabc123 --> abc123

	// TrimPrefix удаляет префикс у строки
	return strings.TrimPrefix(hexStr, "0x")
}
