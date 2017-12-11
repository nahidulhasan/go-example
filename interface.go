package main

import (
	"fmt"
)

type Printer interface {
	Print(string)
}

type Terminal struct {
}

func (t *Terminal) Print(message string) {
	fmt.Println(message)
}

func main() {
	var terminal Terminal
	terminal = Terminal{}
	terminal.Print("Hello!")
}
