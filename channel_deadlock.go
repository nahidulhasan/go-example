package main

func main() {
	c := make(chan int)
	c <- 42    // write to a channel
	val := <-c // read from a channel
	println(val)
}
