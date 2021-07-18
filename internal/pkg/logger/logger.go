package logger

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type LogDetail struct {
	Details interface{} `json:"details,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type structuredLog struct {
	Fn       string     `json:"fn"`
	Metadata *LogDetail `json:"metadata,omitempty"`
}

func Info(fn string, message string, metadata *LogDetail) {
	newLogEntry(fn, metadata).Info(message)
}

func Warn(fn string, message string, metadata *LogDetail) {
	newLogEntry(fn, metadata).Warn(message)
}

func Error(fn string, message string, metadata *LogDetail) {
	newLogEntry(fn, metadata).Error(message)
}

func Debug(fn string, message string, metadata *LogDetail) {
	newLogEntry(fn, metadata).Debug(message)
}

func newLogEntry(fn string, metadata *LogDetail) *logrus.Entry {
	log := structuredLog{
		Fn:       fn,
		Metadata: metadata,
	}

	return log.toStructuredLog()
}

func (sl structuredLog) toStructuredLog() *logrus.Entry {
	var sLog logrus.Fields

	str, _ := json.Marshal(sl)
	json.Unmarshal(str, &sLog)

	return logrus.WithFields(sLog)
}
