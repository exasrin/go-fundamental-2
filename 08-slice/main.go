package main

import "fmt"

func main() {
	// inisialisasi slice
	// bedanya dengan array kita tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)
	newFruits := fruits[0:2]
	fmt.Println(newFruits)

	// fungsi len() dan cap()
	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits2)) // len: 4
	fmt.Println(cap(fruits2)) // cap: 4

	var aFruits2 = fruits2[0:3]
	fmt.Println(len(aFruits2)) // len: 3
	fmt.Println(cap(aFruits2)) // cap: 4

	var bFruits2 = fruits2[1:4]
	fmt.Println(len(bFruits2)) // len: 3
	fmt.Println(cap(bFruits2)) // cap: 3

	// fungsi append()
	var fruits3 = []string{"apel", "mangga"}
	fmt.Println(fruits3)
	fruits3 = append(fruits3, "durian")
	fmt.Println(fruits3)

	var fruits4 = []string{"apple", "grape", "banana"}
	var bFruits4 = fruits4[0:2]

	fmt.Println(cap(bFruits4)) // 3
	fmt.Println(len(bFruits4)) // 2

	fmt.Println(fruits4)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits4) // ["apple", "grape"]

	var cFruits4 = append(bFruits4, "papaya")

	fmt.Println(fruits4)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits4) // ["apple", "grape"]
	fmt.Println(cFruits4) // ["apple", "grape", "papaya"]

	// Fungsi copy()
	var dst = make([]string, 2)
	var src = []string{"apel", "mangga", "durian"}
	n := copy(dst, src)
	fmt.Println(n)
}
