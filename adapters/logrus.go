package adapters

import (
	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

type LogrusAdapter struct {
	entry *logrus.Entry
}

func NewLogrusAdapter() *LogrusAdapter {
	fields := make(logrus.Fields)

	return &LogrusAdapter{entry: logrus.WithFields(fields)}
}

func (this *LogrusAdapter) IsDebugEnabled() bool {
	return this.entry.Logger.IsLevelEnabled(logrus.DebugLevel)
}

func (this *LogrusAdapter) IsTraceEnabled() bool {
	return this.entry.Logger.IsLevelEnabled(logrus.TraceLevel)
}

func (this *LogrusAdapter) WithFields(fields Fields) *LogrusAdapter {
	logrusFields := make(logrus.Fields, len(fields))
	for k, v := range fields {
		logrusFields[k] = v
	}

	return &LogrusAdapter{entry: logrus.WithFields(logrusFields)}
}

func (this *LogrusAdapter) Debugf(format string, args ...interface{}) {
	this.entry.Debugf(format, args...)
}

func (this *LogrusAdapter) Tracef(format string, args ...interface{}) {
	this.entry.Tracef(format, args...)
}
