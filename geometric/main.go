package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	a := make([]int, n)
	for i, part := range parts {
		a[i], _ = strconv.Atoi(part)
	}

	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		square := a[i] * a[i]
		cnt[square] = cnt[square] + 1
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			product := a[i] * a[j]
			ans += cnt[product]
		}
	}

	fmt.Println(ans)
}
