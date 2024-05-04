package main

import "fmt"

func main() {
	const pi = 3.14
	fmt.Println(pi)

	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)
	fmt.Printf("%s, %t, %d, %.2f\n", square, isToday, numeric, floatNum)

	const (
		a = "konstanta"
		b
	)

	const (
		today string = "senin"
		sekarang
		isToday2 = true
	)

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

}
