package main

import (
	"testing"

	utils "github.com/ed-henrique/123-go/utils"
)

func TestFormatTime(t *testing.T) {
	t.Run("formatting time lower than 1 second", func(t *testing.T) {
		got := utils.FormatTime(100)
		want := "0.100s"
		assertCorrectMessage(t, got, want)
	})

	t.Run("formatting time equal to 1 second", func(t *testing.T) {
		got := utils.FormatTime(1000)
		want := "1.000s"
		assertCorrectMessage(t, got, want)
	})

	t.Run("formatting time lower than 1 minute", func(t *testing.T) {
		got := utils.FormatTime(59967)
		want := "59.967s"
		assertCorrectMessage(t, got, want)
	})

	t.Run("formatting time equal to 1 minute", func(t *testing.T) {
		got := utils.FormatTime(60000)
		want := "1:00.000s"
		assertCorrectMessage(t, got, want)
	})

	t.Run("formatting time lower than 1 hour", func(t *testing.T) {
		got := utils.FormatTime(63400)
		want := "1:03.400s"
		assertCorrectMessage(t, got, want)
	})

	t.Run("formatting time equal to 1 hour", func(t *testing.T) {
		got := utils.FormatTime(3600000)
		want := "1:00:00.000s"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
