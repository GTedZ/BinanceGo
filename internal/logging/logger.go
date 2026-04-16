package logging

type Logger interface {
	DEBUG(message string, err ...error)
	DEBUGf(format string, args ...any)
	INFO(message string, err ...error)
	INFOf(format string, args ...any)
	WARN(message string, err ...error)
	WARNf(format string, args ...any)
	ERROR(message string, err ...error)
	ERRORf(format string, args ...any)
}

type LogLevel int

const (
	NONE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case NONE:
		return "NONE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"

	default:
		return "N/A"
	}
}

func (l LogLevel) Color() string {
	switch l {
	case NONE:
		return ""
	case DEBUG:
		return "\x1b[90m"
	case INFO:
		return "\x1b[37m"
	case WARN:
		return "\x1b[33m"
	case ERROR:
		return "\x1b[31m"

	default:
		return ""
	}
}
