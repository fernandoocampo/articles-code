package logging

// Fields define fields type to add in logs.
type Fields map[string]interface{}

// LogFormat log output format
type LogFormat int

const (
	// JSON log output format
	JSON LogFormat = 1 << iota
	// Text log output format
	Text
)

// Level log level type
type Level int

const (
	// Debug log level
	Debug Level = 1 << iota
	// Info log level
	Info
	// Warn log level
	Warn
	// Error log level
	Error
)

// LoggerType log mechanism to use in the application.
type LoggerType int

const (
	// Basic is the logger that uses native logger.
	Basic LoggerType = 1 << iota
	// Logrus is the logger that uses logrus logger.
	Logrus
)

// artifactField is the name of the artifact field in the logs.
const artifactField = "artifact"

// Logger defines behavior for logger mechanism.
type Logger interface {
	// Debug log with debug level
	Debug(message string, fields Fields)
	// Info log with info level
	Info(message string, fields Fields)
	// Warn log with warn level
	Warn(message string, fields Fields)
	// Error log with error level
	Error(message string, fields Fields)
}

// New creates a new logger
func New(selectedLogger LoggerType, artifactName string, level Level) Logger {
	switch selectedLogger {
	case Basic:
		return NewBasicLoggerWithStdout(artifactName, level)
	case Logrus:
		return NewLogrusLoggerWithStdout(artifactName, level)
	default:
		return NewBasicLoggerWithStdout(artifactName, level)
	}
}
