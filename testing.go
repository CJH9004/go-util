package util

import "testing"

func FatalIfErr(tb testing.TB, err error) {
	if err != nil {
		tb.Fatal(NewFileLine(2), err)
	}
}

func Fatal(tb testing.TB, fatal bool, msg string) {
	if fatal {
		tb.Fatal(NewFileLine(2), msg)
	}
}
