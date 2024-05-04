package main

import "fmt"

func main() {
	isRun := true

	if isRun {
		fmt.Println("This is block if")
	}

	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// Variabel Temporary Pada if - else
	var point2 = 8840.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var point4 = 6

	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	/*
		Keyword fallthrough digunakan untuk memaksa proses pengecekan tetap diteruskan ke case selanjutnya
		dengan tanpa menghiraukan nilai kondisinya, efeknya adalah case di pengecekan selanjutnya selalu
		dianggap true (meskipun aslinya bisa saja kondisi tersebut tidak terpenuhi, akan tetap dianggap true).
	*/
	fmt.Println("=======================")
	var point5 = 6

	switch {
	case point5 == 8:
		fmt.Println("perfect")
	case (point5 < 8) && (point5 > 3):
		fmt.Println("awesome")
		fallthrough
	case point5 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
