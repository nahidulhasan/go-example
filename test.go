package main

import  "fmt"

import "strings"

func duplicateCheck(s1 string) int {

	counter := make(map[string]int)

	for _, s := range s1 {
		sl := strings.ToLower(string(s))
		counter[sl] += 1
	}

	fmt.Println("map:", counter)

	sum := 0
	for _, v := range counter {
		if v > 1 {
			sum += 1
		}
	}

	return sum
}

func main() {

	result := duplicateCheck("abbdee" )

	fmt.Println(result)


}
