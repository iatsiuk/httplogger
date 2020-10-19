package httplogger

type Fields map[string]interface{}

type GenericLogger interface {
	IsDebugEnabled() bool
	IsTraceEnabled() bool

	WithFields(fields Fields) GenericLogger

	Debugf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
}

func NewHTTPLogger(logger GenericLogger) string {
	return "test"
}
