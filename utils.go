package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Text colors
const (
	Black int = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// PrintInColor prints text in color
func PrintInColor(message string, color int) {
	fmt.Printf("\033[%dm%s\033[0m", color, message)
}

// PrintInColorln prints text in color
func PrintInColorln(message string, color int) {
	fmt.Printf("\033[%dm%s\033[0m\n", color, message)
}

// LogInColor logs text in color
func LogInColor(message string, color int) {
	log.Printf("\033[%dm%s\033[0m\n\n", color, message)
}

// ExitWithError do `exit 1` logginig error in red
func ExitWithError(message string) {
	LogInColor(message, Red)
	flag.Usage()
	os.Exit(1)
}
