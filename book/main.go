package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inputs := scanner.Text()
	counts := strings.Fields(inputs)

	wallet, _ := strconv.Atoi(counts[1])

	scanner.Scan()
	secondLine := scanner.Text()
	bookStr := strings.Fields(secondLine)

	books := make([]string, len(bookStr))
	copy(books, bookStr)

	sort.Strings(books)

	purchased := make(map[string]bool)
	for _, price := range books {
		priceInt, _ := strconv.Atoi(price)
		if priceInt <= wallet {
			wallet -= priceInt
			purchased[price] = true
		}
	}

	var indexes []int
	for i, price := range bookStr {
		_, ok := purchased[price]
		if ok {
			indexes = append(indexes, i+1)
		}
	}

	fmt.Println(len(purchased))
	for _, v := range indexes {
		fmt.Print(v, " ")
	}
	fmt.Println()

}
