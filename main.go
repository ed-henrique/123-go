package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	now := time.Now()

	go func() {
		for {
			fmt.Printf("\r%v", time.Since(now).Milliseconds())
		}

	}()

	fmt.Scanln(&i)
	timeElapsed := time.Since(now)
	fmt.Println(timeElapsed)
}
