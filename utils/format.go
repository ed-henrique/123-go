package utils

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func PrintCentered(text string) {
	width, _, err := term.GetSize(0)

	if err != nil {
		fmt.Println("Could not print time!")
		os.Exit(1)
	}

	centerHorizontal := (width - len(text)) / 2

	fmt.Printf("\r%s%s", strings.Repeat(" ", centerHorizontal), text)
}

func PrintTime(time int64) {
	width, _, err := term.GetSize(0)

	if err != nil {
		fmt.Println("Could not print time!")
		os.Exit(1)
	}

	formattedTime := formatTime(time)
	centerHorizontal := (width - len(formattedTime)) / 2

	fmt.Printf("\r%s%v", strings.Repeat(" ", centerHorizontal), formattedTime)
}

func formatTime(milliseconds int64) (result string) {
	switch {
	case milliseconds < 1000:
		result = fmt.Sprintf("0.%ds", milliseconds)
	case milliseconds < 10000:
		seconds := milliseconds / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%d.%03ds", seconds, milliseconds)
	case milliseconds < 60000:
		seconds := milliseconds / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%2d.%03ds", seconds, milliseconds)
	case milliseconds < 600000:
		minutes := milliseconds / 60000
		seconds := (milliseconds % 60000) / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%d:%02d.%03ds", minutes, seconds, milliseconds)
	case milliseconds < 3600000:
		minutes := milliseconds / 60000
		seconds := (milliseconds % 60000) / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%2d:%02d.%03ds", minutes, seconds, milliseconds)
	default:
		hours := milliseconds / 3600000
		minutes := (milliseconds % 3600000) / 60000
		seconds := (milliseconds % 60000) / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%d:%02d:%02d.%03ds", hours, minutes, seconds, milliseconds)
	}

	return
}
