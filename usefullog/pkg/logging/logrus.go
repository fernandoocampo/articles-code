package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// logrusLogger contains a logrus logger.
type logrusLogger struct {
	// level log level
	level Level
	// artifact is the name of the artifact that print the log.
	artifact string
	// logger is the logrus logger.
	logger *logrus.Logger
}

// NewLogrusLoggerWithStdout create a logrus logger with stdout output.
func NewLogrusLoggerWithStdout(artifactName string, level Level) Logger {
	return NewLogrusLogger(artifactName, level, os.Stdout)
}

// NewLogrusLogger create a logrus logger.
func NewLogrusLogger(artifactName string, level Level, output io.Writer) Logger {
	logger := logrus.New()
	logger.SetOutput(output)
	logger.SetLevel(getLogrusLevel(level))
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger = logger.WithField(artifactField, artifactName)
	return &logrusLogger{
		level:    level,
		logger:   logger,
		artifact: artifactName,
	}
}

func (l *logrusLogger) Debug(message string, fields Fields) {
	if l.level&Debug != 0 {
		l.logger.WithField(artifactField, l.artifact).
			WithFields(logrus.Fields(fields)).
			Debug(message)
	}
}

func (l *logrusLogger) Info(message string, fields Fields) {
	if l.level&(Debug|Info) != 0 {
		l.logger.WithField(artifactField, l.artifact).
			WithFields(logrus.Fields(fields)).
			Info(message)
	}
}

func (l *logrusLogger) Warn(message string, fields Fields) {
	if l.level&(Debug|Info|Warn) != 0 {
		l.logger.WithField(artifactField, l.artifact).
			WithFields(logrus.Fields(fields)).
			Warn(message)
	}
}

func (l *logrusLogger) Error(message string, fields Fields) {
	if l.level&(Debug|Info|Warn|Error) != 0 {
		l.logger.WithField(artifactField, l.artifact).
			WithFields(logrus.Fields(fields)).
			Error(message)
	}
}

// getLogrusLevel get the logrus level based on allowed levels.
func getLogrusLevel(level Level) logrus.Level {
	switch level {
	case Debug:
		return logrus.DebugLevel
	case Info:
		return logrus.InfoLevel
	case Warn:
		return logrus.WarnLevel
	case Error:
		return logrus.ErrorLevel
	default:
		return logrus.DebugLevel
	}
}
