package logging

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// prefixFormat constains the format for the prefix where
// the first parameter is the level and the second the artifact.
const prefixFormat = "%s [%s]: "

// prefixFunc build the prefix name.
var prefixFunc = func(levelName, artifactName string) string {
	return fmt.Sprintf(prefixFormat, levelName, artifactName)
}

// basicLogger use log library to log.
type basicLogger struct {
	// artifact is the name of the artifact that print the log.
	artifact string
	// level log level
	level       Level
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

// NewBasicLoggerWithStdout create a logger with stout output.
func NewBasicLoggerWithStdout(artifactName string, level Level) Logger {
	return NewBasicLogger(artifactName, level, os.Stdout)
}

// NewBasicLogger creates a new basic logger.
func NewBasicLogger(artifactName string, level Level, output io.Writer) Logger {
	return &basicLogger{
		level:       level,
		artifact:    artifactName,
		debugLogger: log.New(output, "DEBUG: ", log.Ldate|log.Ltime),
		infoLogger:  log.New(output, "INFO: ", log.Ldate|log.Ltime),
		warnLogger:  log.New(output, "WARN: ", log.Ldate|log.Ltime),
		errorLogger: log.New(output, "ERROR: ", log.Ldate|log.Ltime),
	}
}

func (b *basicLogger) Debug(message string, fields Fields) {
	if b.level&Debug != 0 {
		b.debugLogger.Println(b.formatLogText(message, fields))
	}

}
func (b *basicLogger) Info(message string, fields Fields) {
	if b.level&(Debug|Info) != 0 {
		b.infoLogger.Println(b.formatLogText(message, fields))
	}

}
func (b *basicLogger) Warn(message string, fields Fields) {
	if b.level&(Debug|Info|Warn) != 0 {
		b.warnLogger.Println(b.formatLogText(message, fields))
	}

}
func (b *basicLogger) Error(message string, fields Fields) {
	if b.level&(Debug|Info|Warn|Error) != 0 {
		b.errorLogger.Println(b.formatLogText(message, fields))
	}
}

// formatLogText concat given message and fields in a string message.
func (b *basicLogger) formatLogText(message string, fields Fields) string {
	var buffer bytes.Buffer
	buffer.WriteString("msg=")
	buffer.WriteString(message)
	buffer.WriteString(" artifact=")
	buffer.WriteString(b.artifact)
	for k, v := range fields {
		buffer.WriteString(fmt.Sprintf(" %s=%v", k, v))
	}
	return buffer.String()
}
