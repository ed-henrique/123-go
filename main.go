package main

import (
	"fmt"
	"time"

	utils "123-go/utils"
)

func main() {
	var i int
	now := time.Now()

	go func() {
		for {
			fmt.Printf("\r%v", utils.FormatTime(time.Since(now).Milliseconds()))
		}
	}()

	fmt.Scanln(&i)
	timeElapsed := time.Since(now)
	fmt.Println("\rTime elapsed:", timeElapsed.String())
}
