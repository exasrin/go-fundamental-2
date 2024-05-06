package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	var numbers2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	// Immediately-Invoked Function Expression (IIFE)
	var newNumbers2 = func(min int) []int {
		var r []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers2)
	fmt.Println("filtered number :", newNumbers2)

	var max2 = 3
	var numbers3 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers3 = findMax(numbers3, max2)
	var theNumbers3 = getNumbers3()

	fmt.Println("numbers3\t:", numbers3)
	fmt.Printf("find \t: %d\n\n", max2)

	fmt.Println("found \t:", howMany)     // 9
	fmt.Println("value \t:", theNumbers3) // [2 3 0 3 2 0 2 0 3]
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
