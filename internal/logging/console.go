package logging

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|0x0004)
}

func print(level LogLevel, message string) {
	fmt.Print(level.Color() + message + "\x1b[0m")
}
