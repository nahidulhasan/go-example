package main

func main() {
	c := make(chan int)

	// Make the writing operation be performed in
	// another goroutine.
	go func() {
		c <- 42
	}()
	val := <-c
	println(val)
}
