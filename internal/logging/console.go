package logging

import (
	"fmt"
)

func print(_ LogLevel, message string) {
	fmt.Print(message)
}
