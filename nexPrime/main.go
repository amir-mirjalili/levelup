package main

import (
	"fmt"
	"math"
)

func main() {
	var number int

	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)
	}

	for i := number + 1; i <= int(math.Pow10(9)); i++ {
		if primeCheck(i) {
			fmt.Println(i)
			break
		}
	}

}

/*
1.  Start with 2: If the number is less than 2, it’s not prime.
2.  Try dividing the number by all integers from 2 up to the square root of the number.
3.  If none divide evenly, it’s prime. If any do, it’s not.
*/
func primeCheck(n int) bool {

	if n < 2 {
		return false
	}

	sqrN := math.Sqrt(float64(n))
	for i := 2; i <= int(sqrN); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
