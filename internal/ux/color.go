package ux

import "fmt"

const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[32m"
	ColorYellow  = "\x1b[33m"
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
