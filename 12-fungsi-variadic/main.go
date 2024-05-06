package main

import "fmt"

func main() {
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate(numbers...)
	// var avg = calculate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	total := 0
	for _, v := range numbers {
		total += v
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}
