package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countXORTriplets(a []int) int {
	n := len(a)
	var answer int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if (a[i] ^ a[j]) == a[k] {
					answer++
				}
			}
		}
	}

	return answer
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	lengthStr := scanner.Text()
	n, err := strconv.Atoi(lengthStr)
	if err != nil {
		fmt.Printf("Error parsing length: %s\n", lengthStr)
		return
	}

	scanner.Scan()
	input := scanner.Text()

	parts := strings.Fields(input)

	if len(parts) != n {
		fmt.Printf("Expected %d elements, but got %d\n", n, len(parts))
		return
	}

	var a []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error parsing number: %s\n", part)
			return
		}
		a = append(a, num)
	}

	result := countXORTriplets(a)
	fmt.Println(result)
}
