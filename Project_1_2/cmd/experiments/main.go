package main

import "fmt"

func main() {
	const language = "Go"
	var name string = "ЕгоЗа"
	completed := 1
	var ready bool = true
	var hours float64 = 3.5

	fmt.Println("Изучаем", language)
	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Выполнено: %d\n", completed)
	fmt.Printf("Готов продолжать: %t\n", ready)
	fmt.Printf("Продолжительность: %.1f часа\n", hours)

	completed = 2
	fmt.Printf("После изменения: %d\n", completed)

	var pages int
	fmt.Print("Введите число страниц: ")
	_, err := fmt.Scan(&pages)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if pages < 0 {
		fmt.Println("Ошибка: число не должно быть отрицательным.")
		return
	}
	fmt.Printf("Будет обработано страниц: %d\n", pages)
}
