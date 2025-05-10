package colorset

import (
	"fmt"
	"math/rand"
)

// ColorRand generates and returns a random ANSI foreground color escape code
// for basic terminal colors. The returned string can be used to set the text
// color in compatible terminals.
//
// Supported colors correspond to codes 30 through 37:
//
//   30: Black
//   31: Red
//   32: Green
//   33: Yellow
//   34: Blue
//   35: Magenta
//   36: Cyan
//   37: White
//
// Example usage:
//
//	fmt.Printf("%sHello colorful world!%s\n", colorset.ColorRand(), colorset.ColorReset())
func ColorRand() string {
	return fmt.Sprintf("\033[0%dm", rand.Intn(8)+30) // Generates a number from 30 to 37
}

// ColorSet returns a specific ANSI foreground color escape code based on the
// provided short color name. Useful when a deterministic color is needed.
//
// Supported colors:
//
//	"r" - Red
//	"g" - Green
//	"y" - Yellow
//	"b" - Blue
//	"p" - Magenta (often seen as purple)
//	"c" - Cyan
//	"w" - White
//	"_" - Reset (same as ColorReset)
//
// Example usage:
//
//	fmt.Print(colorset.ColorSet("g"), "This is green text", colorset.ColorReset())
func ColorSet(color string) string {
	colors := map[string]string{
		"r":   "\033[031m",
		"g":   "\033[032m",
		"y":   "\033[033m",
		"b":   "\033[034m",
		"p":   "\033[035m",
		"c":   "\033[036m",
		"w":   "\033[037m",
		"_": "\033[0m",
	}

	return colors[color] // Directly return the color code
}

// ColorReset returns the ANSI escape sequence to reset all terminal text
// formatting including color. Use this after applying any color to avoid
// unintended style bleed.
//
// This is equivalent to "\x1b[0m".
//
// Example usage:
//
// fmt.Printf("%sImportant message%s\n", colorset.ColorSet("r"), colorset.ColorReset())
func ColorReset() string {
	return "\033[0m" // Reset color code
}
