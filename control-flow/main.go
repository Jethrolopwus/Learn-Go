package main

import "fmt"

func main() {

	// ============ MAP ==============

	users := map[string]int{
		"Jah":   1,
		"Bob":   23,
		"ALice": 3,
		"Brad":  40,
		"Jane":  55,
	}

	for key, value := range users {
		fmt.Printf("key %s value %d\n ", key, value)
	}

	// ========= using a slice ===========

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(numbers); i++ {
		// fmt.Println("it ", i)
		fmt.Println(numbers[i])
	}
	// definition
	for i := 0; i < 10; i++ {
		fmt.Println("it ", i)
	}

	// ======RANGE====
	cars := []string{"Benz", "Honda", "Toyota", "BMW"}
	for _, car := range cars {
		fmt.Println(car)
	}

	// ====== BREAK ========

	names := []string{"Joe", "Lang", "Jane", "Brad"}
	for _, name := range names {

		if name == "Joe" {
			break
		}

	}
	fmt.Println("Break out of the Loop")
}
