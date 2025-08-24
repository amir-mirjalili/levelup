package main

import (
	"fmt"
	"strconv"
)

var lucky []int

func generateLuckyNumbers() {
	for k := 1; k < 10; k++ {
		for mask := 0; mask < (1 << k); mask++ {
			a := ""
			for i := 0; i < k; i++ {
				if (mask & (1 << i)) != 0 {
					a += "7"
				} else {
					a += "4"
				}
			}
			num, _ := strconv.Atoi(a)
			lucky = append(lucky, num)
		}
	}
}

func solve(n int) {
	ans := 4
	for _, m := range lucky {
		if m <= n && m > ans {
			ans = m
		}
	}
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)

	generateLuckyNumbers()

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		solve(n)
	}
}
