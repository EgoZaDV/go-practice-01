package main

import "fmt"

func main() {
	var minutes int
	fmt.Print("Введите целое неотрицательное количество минут: ")
	_, err := fmt.Scan(&minutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}
	if minutes < 0 {
		fmt.Println("Ошибка: число не должно быть отрицательным")
		return
	}

	const mins_per_hour = 60
	var hours int = minutes / mins_per_hour
	mins_left := minutes % mins_per_hour

	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d\n", mins_left)
}
