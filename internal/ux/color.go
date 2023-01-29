package ux

import "fmt"

const (
	ColorGreen   = "\x1b[32m"
	ColorYellow  = "\x1b[33m"
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorBlue    = "\x1b[94m"
)

func green(str string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, str, ColorDefault)
}

func red(str string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, str, ColorDefault)
}

func yellow(str string) string {
	return fmt.Sprintf("%s%s%s", ColorYellow, str, ColorDefault)
}

func blue(str string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, str, ColorDefault)
}
