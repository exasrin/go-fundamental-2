package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("ini indeks ke-%d\n", i)
	}

	idx := 0
	for idx < 5 {
		fmt.Printf("lopp-%d\n", idx)
		idx++
	}

	var i = 0

	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}

	var xs = "123"
	for i, v := range xs {
		fmt.Println("indeks:", i, "value:", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50}
	for _, v := range ys {
		fmt.Println("value:", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2}
	for k, v := range kvs {
		fmt.Println("key:", k, "value:", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Println(i) // 01234
	}
	fmt.Println("=================")
	fmt.Println(len(ys))

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
