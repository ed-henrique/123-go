package main

import (
	t "github.com/ed-henrique/123-go/pkg"
	u "github.com/ed-henrique/123-go/utils"
)

func main() {
	duration := u.GetFlags()

	switch {
	case duration > 0:
		t.Timer(duration)
	default:
		t.Stopwatch()
	}
}
