package main

import "fmt"

func main() {
	const course = "Бэкенд-разработка на языке Go"
	var name string = "Комаеда"
	var group string = "Учебная-02"
	finished := 4
	interest := true

	fmt.Println("Участник:", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Println("Курс:", course)
	fmt.Printf("Завершено работ: %d\n", finished)
	fmt.Printf("Интерес к бэкенду: %t\n", interest)
}
