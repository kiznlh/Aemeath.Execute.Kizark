package utils

import (
	"fmt"
	"time"

	"github.com/muesli/termenv"
)

// Print text with duration
func SlowPrint(text termenv.Style, delay time.Duration) {
	for _, char := range text.String() {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}

// Println text with duration
func SlowPrintln(text termenv.Style, delay time.Duration) {
	SlowPrint(text, delay)
	fmt.Println()
}
