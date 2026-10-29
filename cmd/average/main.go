package main

import "fmt"

func main() {
	var n1 int = 20
	var n2 int = 25
	var n3 int = 35

	var summa int = n1 + n2 + n3
	fmt.Println("Сумма:", summa)
	var fsumma float64 = float64(summa)
	var nums float64 = 3.0
	var avg float64 = fsumma / nums
	fmt.Printf("Среднее арифметическое: %.2f\n", avg)
}
