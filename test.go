package main

import  "fmt"

import s "strings"



func FizzBuzzCuckooClockTest(time string) string {

	arr := s.Split(time, "-")

	fmt.Println(arr[0])
	fmt.Println(arr[1])

	return arr[0]

}

func main() {

	result := FizzBuzzCuckooClockTest("11:15" )

	fmt.Println(result)


}
