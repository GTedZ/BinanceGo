package logging

import (
	"fmt"
	"os"
	"time"
)

type binanceLogger struct {
	filename      string
	logLevel      LogLevel
	printLogLevel LogLevel
}

func nowFormatted() string {
	return time.Now().Format("2006/01/02 15:04:05.000")
}

func (bl *binanceLogger) log(level LogLevel, message string) {
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

func (bl *binanceLogger) DEBUG(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(DEBUG, message)
}

func (bl *binanceLogger) DEBUGf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	bl.log(DEBUG, message)
}

func (bl *binanceLogger) INFO(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(INFO, message)
}

func (bl *binanceLogger) INFOf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	bl.log(INFO, message)
}

func (bl *binanceLogger) WARN(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(WARN, message)
}

func (bl *binanceLogger) WARNf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	bl.log(WARN, message)
}

func (bl *binanceLogger) ERROR(message string, err ...error) {
	for i, e := range err {
		message += fmt.Sprintf("\n\t> err %d: %s", i, e.Error())
	}
	bl.log(ERROR, message)
}

func (bl *binanceLogger) ERRORf(format string, args ...any) {
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

	if config.LogLevel != NONE {
		_, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("[BinanceLogger] Error opening file:", err)
			return nil, err
		}
	}

	return &binanceLogger{
		filename:      filename,
		logLevel:      config.LogLevel,
		printLogLevel: config.PrintLogLevel,
	}, nil
}
