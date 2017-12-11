package main

import (
	"strings"
	"fmt"
)

func duplicateCount(s1 string) int {

	counter := make(map[string]int)

	for _, s := range s1 {
		sl := strings.ToLower(string(s))
		counter[sl] += 1
	}

	sum := 0
	for _, v := range counter {
		if v > 1 {
			sum += 1
		}
	}

	return sum
}

func main() {

	result := duplicateCount("abbde" )

	fmt.Println(result)


}

