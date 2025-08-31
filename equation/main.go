package main

import (
	"fmt"
	"math"
)

func solve() {
	var n, m int
	fmt.Scan(&n, &m)

	N := int(math.Sqrt(float64(n+m)/2)) + 2
	ans := 0

	for a := -N; a <= N; a++ {
		for b := -N; b <= N; b++ {
			if a*a+a*b+b*b == n && a*a-a*b+b*b == m {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		solve()
		t--
	}
}
