package main

import  "fmt"

import s "strings"

import  "strconv"

func FizzBuzzCuckooClock(time string) string {

	var str string

	arr := s.Split(time, ":")

	hour, _ := strconv.Atoi(arr[0])

	minute, _ := strconv.Atoi(arr[1])


	if minute == 30 {

		str = "Cuckoo"

	} else if hour == 0 && minute == 0  {

		str = s.TrimSpace(s.Repeat("Cuckoo ", 12))

	} else if minute == 0 && hour != 0 {

		var num int

		if hour != 12 {
			num = hour % 12
		} else {
			num = 12
		}

		str = s.TrimSpace(s.Repeat("Cuckoo ", num))

	} else if (minute % 3 == 0) && (minute % 5 == 0){

		str = "Fizz Buzz"

	} else if minute % 3 == 0 {

		str = "Fizz"

	} else if minute % 5 == 0 {

		str = "Buzz"

	} else {

		str = "tick"

	}

	return str
}

func main() {

	result := FizzBuzzCuckooClock("11:15")

	fmt.Println(result)

}


