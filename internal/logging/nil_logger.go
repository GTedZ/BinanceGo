package logging

type NilLogger struct{}

////
// Interface Implementations
////

func (bl *NilLogger) DEBUG(message string, err ...error) {}

func (bl *NilLogger) DEBUGf(format string, args ...any) {}

func (bl *NilLogger) INFO(message string, err ...error) {}

func (bl *NilLogger) INFOf(format string, args ...any) {}

func (bl *NilLogger) WARN(message string, err ...error) {}

func (bl *NilLogger) WARNf(format string, args ...any) {}

func (bl *NilLogger) ERROR(message string, err ...error) {}

func (bl *NilLogger) ERRORf(format string, args ...any) {}

////
// Public Export
////

func NewNilLogger() Logger {
	return &NilLogger{}
}
