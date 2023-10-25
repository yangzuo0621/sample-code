package main

// https://en.wikipedia.org/wiki/Braille_Patterns

import (
	"fmt"
	"os"
	"time"
)

func main() {
	result := supervisor()
	fmt.Println("Answer:", result)
}

// https://stackoverflow.com/questions/15433188/what-is-the-difference-between-r-n-r-and-n
// All 3 of them represent the end of a line. But...
//
// + \r (Carriage Return) → moves the cursor to the beginning of the line without advancing to the next line
// + \n (Line Feed) → moves the cursor down to the next line without returning to the beginning of the line
//                    — In a *nix environment \n moves to the beginning of the line.
// + \r\n (End Of Line) → a combination of \r and \n

func spin(msg string, done <-chan struct{}) {
	fmt.Fprintf(os.Stdout, "%s\n", "spinner goroutine: spinning!")
	chars := []byte{'\\', '|', '/', '-'}
	index := 0
	for {
		select {
		case <-done:
			blanks := make([]byte, len(msg))
			os.Stdout.Write([]byte(fmt.Sprintf("\r%s\r", string(blanks)))) // clear line
			// os.Stdout.Sync()
			// fmt.Fprintf(os.Stdout)
			return
		default:
			status := fmt.Sprintf("\r%c %s", chars[index], msg) // spin in oneline
			index = (index + 1) % len(chars)
			fmt.Fprintf(os.Stdout, "%s", status)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func slow() int {
	time.Sleep(3 * time.Second)
	return 42
}

func supervisor() int {
	done := make(chan struct{})
	go spin("thinking!", done)
	fmt.Fprintf(os.Stdout, "%s\n", "main goroutine: spinning!")
	result := slow()
	done <- struct{}{}
	return result
}
