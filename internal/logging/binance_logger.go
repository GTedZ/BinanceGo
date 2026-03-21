package logging

import (
	"fmt"
	"os"
	"time"
)

type BinanceLogger struct {
	filename      string
	logLevel      LogLevel
	printLogLevel LogLevel
}

func nowFormatted() string {
	return time.Now().Format("2006/01/02 15:04:05.000")
}

func (bl *BinanceLogger) log(level LogLevel, message string) {
	timeStr := nowFormatted()

	formattedMessage := fmt.Sprintf("%s: [%s] %s", timeStr, level.String(), message)

	if bl.logLevel != NONE && level >= bl.logLevel {
		err := appendFile(bl.filename, formattedMessage)
		if err != nil {
			fmt.Printf("[BinanceLogger] Failed to write log: %s", err.Error())
		}
	}

	if bl.printLogLevel != NONE && level >= bl.printLogLevel {
		print(level, formattedMessage)
	}
}

////
// Interface Implementations
////

func (bl *BinanceLogger) DEBUG(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(DEBUG, message)
}

func (bl *BinanceLogger) DEBUGf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	bl.log(DEBUG, message)
}

func (bl *BinanceLogger) INFO(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(INFO, message)
}

func (bl *BinanceLogger) INFOf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	bl.log(INFO, message)
}

func (bl *BinanceLogger) WARN(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(WARN, message)
}

func (bl *BinanceLogger) WARNf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	bl.log(WARN, message)
}

func (bl *BinanceLogger) ERROR(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(ERROR, message)
}

func (bl *BinanceLogger) ERRORf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	bl.log(ERROR, message)
}

////
// Public Exports
////

type Config struct {
	// Filename is the target file that logs will be written to
	// Default "logs.txt"
	Filename string

	// PrintLogLevel sets the log level at which they are logged to `Filename`
	// Default "INFO"
	LogLevel LogLevel

	// PrintLogLevel sets the log level at which they are printed to the console
	// Default "INFO"
	PrintLogLevel LogLevel
}

func New(config Config) (Logger, error) {
	filename := "logs.txt"
	if config.Filename != "" {
		filename = config.Filename
	}

	_, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[BinanceLogger] Error opening file:", err)
		return nil, err
	}

	return &BinanceLogger{
		filename:      filename,
		logLevel:      config.LogLevel,
		printLogLevel: config.PrintLogLevel,
	}, nil
}
