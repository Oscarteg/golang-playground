package src

import (
	"io"
	"log"
	"os"
)

func run() {
	if err := write("readme.txt", "This is a readme"); err != nil {
		log.Fatal()
	}
}

// In Go, it is considered a safe and accepted practice to call Close() more than once without affecting the behavior of your program. If Close() is going to return an error, it will do so the first time it is called. This allows us to call it explicitly in the successful path of execution in our function.

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}

	return file.Close()
}
