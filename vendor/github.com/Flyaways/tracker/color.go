package tracker

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

// NoColor defines if the output is colorized or not. It's dynamically set to
// false or true based on the stdout's file descriptor referring to a terminal
// or not. This is a global option and affects all colors. For more control
// over each color block use the methods DisableColor() individually.
var NoColor = !isatty.IsTerminal(os.Stdout.Fd())

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params  []Attribute
	noColor *bool
}

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// New returns a newly created color object.
func New(value ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0)}
	c.Add(value...)
	return c
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(p ...Attribute) *Color {
	c := New(p...)
	c.Set()
	return c
}

// Unset resets all escape attributes and clears the output. Usually should
// be called after Set().
func Unset() {
	if NoColor {
		return
	}

	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprintf(Output, c.format())
	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

// Output defines the standard output of the print functions. By default
// os.Stdout is used.
var Output = colorable.NewColorableStdout()

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjuction with color.Output.
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// sequence returns a formated SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// wrap wraps the s string with the colors attributes. The string is ready to
// be printed.
func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *Color) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor enables the color output. Use it in conjuction with
// DisableColor(). Otherwise this method has no side effects.
func (c *Color) EnableColor() {
	c.noColor = boolPtr(false)
}

func (c *Color) isNoColorSet() bool {
	// check first if we have user setted action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

func boolPtr(v bool) *bool {
	return &v
}

func printString(format string, p Attribute, a ...interface{}) string {
	if len(a) == 0 {
		a = append(a, format)
		format = "%s"
	}
	return New(p).SprintfFunc()(format, a...)
}

// BlackString is an convenient helper function to return a string with black
// foreground.
func Black(format string, a ...interface{}) string { return printString(format, FgBlack, a...) }

// RedString is an convenient helper function to return a string with red
// foreground.
func Red(format string, a ...interface{}) string { return printString(format, FgRed, a...) }

// GreenString is an convenient helper function to return a string with green
// foreground.
func Green(format string, a ...interface{}) string { return printString(format, FgGreen, a...) }

// YellowString is an convenient helper function to return a string with yellow
// foreground.
func Yellow(format string, a ...interface{}) string { return printString(format, FgYellow, a...) }

// BlueString is an convenient helper function to return a string with blue
// foreground.
func Blue(format string, a ...interface{}) string { return printString(format, FgBlue, a...) }

// MagentaString is an convenient helper function to return a string with magenta
// foreground.
func Magenta(format string, a ...interface{}) string {
	return printString(format, FgMagenta, a...)
}

// CyanString is an convenient helper function to return a string with cyan
// foreground.
func Cyan(format string, a ...interface{}) string { return printString(format, FgCyan, a...) }

// WhiteString is an convenient helper function to return a string with white
// foreground.
func White(format string, a ...interface{}) string { return printString(format, FgWhite, a...) }
