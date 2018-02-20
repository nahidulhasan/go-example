package main

func main() {
	// Create a channel to synchronize goroutines
	done := make(chan bool)

	// Execute println in goroutine
	go func() {
		println("goroutine message")

		// Tell the main function everything is done.
		// This channel is visible inside this goroutine because
		// it is executed in the same address space.
		done <- true
	}()

	println("main function message")
	<-done // Wait for the goroutine to finish
}
