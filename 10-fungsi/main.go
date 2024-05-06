package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var names = []string{"Alfaathir", "Rasyid", "Sulaiman"}
	printMessage("Hello", names)

	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

// Penerapan Fungsi
func printMessage(message string, arr []string) {
	var newString = strings.Join(arr, " ")
	fmt.Println(message, newString)
}

// fungsi dengan return value
func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

// Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
