package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alfaathir"
	names[1] = "Rasyid"
	names[2] = "Sulaiman"
	fmt.Println(names)
	for _, v := range names {
		fmt.Println(v)
	}

	// Inisialisasi Nilai Array Dengan Gaya Vertikal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// array multidimensi
	// var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	var numbers2 = [2][3]int{{1, 3, 5}, {7, 9, 11}}

	//fmt.Println(numbers1)
	fmt.Println(numbers2)

	// looping for array
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruits2); i++ {
		fmt.Println("for fruit", fruits2[i])
	}

	// looping for range array
	for _, v := range fruits2 {
		fmt.Println("for range", v)
	}
	/*
		jika ingin mengakses indeksnya saja
			for i, _ := range fruits { }
		atau
			for i := range fruits { }
	*/

	// Alokasi Elemen Array Menggunakan Keyword make
	var fruits3 = make([]string, 2)
	fruits3[0] = "manggo"
	fruits3[1] = "jamboo"
	fmt.Println(fruits3)
}
