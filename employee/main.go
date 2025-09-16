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
	counts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(counts[0])

	scanner.Scan()
	tasksStr := strings.Fields(scanner.Text())

	var taskArr []int
	for i := 0; i < n; i++ {
		task, _ := strconv.Atoi(tasksStr[i])
		taskArr = append(taskArr, task)
	}
	sort.Ints(taskArr)

	answer := 0
	currentTime := 0
	for i := 0; i < len(taskArr); i++ {
		if taskArr[i] > currentTime {
			answer++
			currentTime++
		}
	}
	fmt.Println(answer)
}
