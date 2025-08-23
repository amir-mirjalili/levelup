package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("NO")
		os.Exit(0)
	}

	if n == 2 || n == 3 {
		fmt.Println("YES")
		os.Exit(0)
	}

	if n%2 == 0 || n%3 == 0 {
		fmt.Println("NO")
		os.Exit(0)
	}

	limit := int(math.Sqrt(float64(n) + 1.1))

	for i := 5; i <= limit; i += 6 {
		if n%i == 0 {
			fmt.Println("NO")
			os.Exit(0)
		}
	}

	for i := 7; i <= int(math.Sqrt(float64(n)*0.5)+1.1); i += 7 {
		if n%i == 0 {
			fmt.Println("NO")
			os.Exit(0)
		}
	}

	fmt.Println("YES")
}
