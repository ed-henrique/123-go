package utils

import "fmt"

func FormatTime(milliseconds int64) (result string) {
	if milliseconds < 1000 {
		result = fmt.Sprintf("0.%ds", milliseconds)
	} else if milliseconds < 10000 {
		seconds := milliseconds / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%d.%03ds", seconds, milliseconds)
	} else if milliseconds < 60000 {
		seconds := milliseconds / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%2d.%03ds", seconds, milliseconds)
	} else if milliseconds < 600000 {
		minutes := milliseconds / 60000
		seconds := (milliseconds % 60000) / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%d:%02d.%03ds", minutes, seconds, milliseconds)
	} else if milliseconds < 3600000 {
		minutes := milliseconds / 60000
		seconds := (milliseconds % 60000) / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%2d:%02d.%03ds", minutes, seconds, milliseconds)
	} else {
		hours := milliseconds / 3600000
		minutes := (milliseconds % 3600000) / 60000
		seconds := (milliseconds % 60000) / 1000
		milliseconds := milliseconds % 1000
		result = fmt.Sprintf("%d:%02d:%02d.%03ds", hours, minutes, seconds, milliseconds)
	}

	return
}
