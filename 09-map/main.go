package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println(chicken)

	// Inisialisasi Nilai Map
	var data map[string]int
	// data["one"] = 1
	// akan muncul error!

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken1)
	fmt.Println(chicken2)

	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// loo for range
	var chicken3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken3 {
		fmt.Println(key, "  \t:", val)
	}

	// Menghapus Item Map
	var chicken4 = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken4)) // 2
	fmt.Println(chicken4)

	delete(chicken4, "januari")

	fmt.Println(len(chicken4)) // 1
	fmt.Println(chicken4)

	// Deteksi Keberadaan Item Dengan Key Tertentu
	var chicken5 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken5["januari"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi Slice & Map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	var chickens2 = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	var datas = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
	fmt.Println("====================================")
	for _, chicken := range chickens2 {
		fmt.Println(chicken["gender"], chicken["name"])
	}
	fmt.Println("====================================")
	for _, data := range datas {
		fmt.Println(data["gender"], data["name"])
	}
}
