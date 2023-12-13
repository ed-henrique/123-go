package pkg

import (
	"fmt"
	"time"

	u "github.com/ed-henrique/123-go/utils"
)

func Stopwatch() {
	var i int
	now := time.Now()

	fmt.Println("\033[2J")

	go func() {
		for {
			u.PrintTime(time.Since(now).Milliseconds())
		}
	}()

	fmt.Scanln(&i)
	fmt.Println("\033[2J")
	u.PrintTime(time.Since(now).Milliseconds())
	fmt.Println()
}
