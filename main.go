package main

import (
	"flag"
	"fmt"
	"time"

	u "github.com/ed-henrique/123-go/utils"
)

var flagTimer int64

func main() {
	flag.Int64Var(&flagTimer, "time", 0, "Starting time to countdown from in milliseconds")
	flag.Parse()

	if flagTimer != 0 {
		startTime := time.Now().Add(time.Millisecond * time.Duration(flagTimer))

		for time.Until(startTime) > 0 {
			fmt.Printf("\x1b[2K\r%v", u.FormatTime(time.Until(startTime).Milliseconds()))
		}

		fmt.Println("\x1b[2K\\rFinished!")
	} else {
		var i int
		now := time.Now()

		go func() {
			for {
				fmt.Printf("\r%v", u.FormatTime(time.Since(now).Milliseconds()))
			}
		}()

		fmt.Scanln(&i)
		timeElapsed := time.Since(now)
		fmt.Println("\rTime elapsed:", u.FormatTime(timeElapsed.Milliseconds()))
	}
}
