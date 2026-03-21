package logging

import (
	"os"
)

func appendFile(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write data to file
	_, err = file.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}
