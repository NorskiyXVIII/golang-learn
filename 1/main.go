// объявление главного модуля(пакета)
package main

// подключение библиотеки fmt
import "fmt"

func main() { // создание главной функции
	fmt.Print("Привет, мир!\n")

	/*var x uint16
	fmt.Println(x)
	x = 4
	fmt.Println(x)*/

	y, x := 3, 4
	fmt.Println("y = ", y, "\nx = ", x)

	g := 'Д' // rune type - utf8 char
	fmt.Print("Д code = ", g, "\n")

}
