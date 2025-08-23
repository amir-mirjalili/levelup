package main

import "fmt"

func main() {
	var number int

	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)
	}

	divisors := 0
	for i := number; i >= 1; i-- {
		if number%i == 0 {
			divisors++
		}
	}

	if divisors == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
