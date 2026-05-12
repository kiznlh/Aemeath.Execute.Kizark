package utils

import (
	"fmt"
	"time"

	"github.com/muesli/termenv"
)

func SlowPrint(text termenv.Style, delay time.Duration) {
	for _, char := range text.String() {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}
