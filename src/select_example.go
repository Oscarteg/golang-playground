package src

import "fmt"

func Fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		// Blocks until one of its cases can run
		// 	Execute that case. If multiple? Choose random
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}
