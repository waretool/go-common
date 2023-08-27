package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
)

type customFormatter struct {
	additionalFields map[string]string
	formatter        logrus.Formatter
}

func (f customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	for k, v := range f.additionalFields {
		entry.Data[k] = v
	}
	byteLog, err := f.formatter.Format(entry)
	// use just 1 space to separate log fields
	trimmedLog := strings.Join(strings.Fields(string(byteLog)), " ") + "\n"
	return []byte(trimmedLog), err
}
