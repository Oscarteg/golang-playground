package src_test

import (
	"fmt"
	"golang-playground/src"
	"testing"
)

func TestFibonacci(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	src.Fibonacci(c, quit)
}
