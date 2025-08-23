package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var builder strings.Builder

	for mask := 0; mask < (1 << n); mask++ {
		builder.Reset()
		builder.WriteByte('{')

		first := true
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if !first {
					builder.WriteString(", ")
				}
				builder.WriteString(strconv.Itoa(i + 1))
				first = false
			}
		}
		builder.WriteString("}\n")

		writer.WriteString(builder.String())
	}
}
