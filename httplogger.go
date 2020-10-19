package httplogger

type Fields map[string]interface{}

type Logger interface {
	IsDebugEnabled() bool
	IsTraceEnabled() bool

	WithFields(fields Fields) *Logger

	Debugf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
}

func NewHTTPLogger(logger Logger) string {
	return "test"
}
