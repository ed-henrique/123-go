package pkg

import (
	"fmt"
	"time"

	u "github.com/ed-henrique/123-go/utils"
)

func Timer(duration int64) {
	fmt.Println("\033[2J")
	startTime := time.Now().Add(time.Millisecond * time.Duration(duration))

	for t := time.Until(startTime); t > 0; t = time.Until(startTime) {
		u.PrintTime(t.Milliseconds())
	}

	fmt.Print("\033[2J")
	u.PrintCentered("Finished!")
	fmt.Println()
}
