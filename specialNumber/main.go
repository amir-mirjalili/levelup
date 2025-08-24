package main

import (
	"fmt"
)

func sumDigits(n int) int {
	result := 0
	for n != 0 {
		result += n % 10
		n /= 10
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	n++
	for sumDigits(n)%7 != 0 {
		n++
	}

	fmt.Println(n)

}
