package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Level represents the severity level of a log message
type Level int

// Log levels
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

// Logger is a simple structured logger
type Logger struct {
	level     Level
	output    io.Writer
	prefix    string
	showCaller bool
}

var defaultLogger *Logger

func init() {
	// Initialize default logger
	defaultLogger = NewLogger(INFO, os.Stdout, "GenRep")
}

// NewLogger creates a new logger with specified level, output and prefix
func NewLogger(level Level, output io.Writer, prefix string) *Logger {
	return &Logger{
		level:     level,
		output:    output,
		prefix:    prefix,
		showCaller: true,
	}
}

// SetLevel sets the logging level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// EnableCallerInfo enables or disables the inclusion of caller information in log messages
func (l *Logger) EnableCallerInfo(enabled bool) {
	l.showCaller = enabled
}

// log logs a message at the specified level
func (l *Logger) log(level Level, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	message := fmt.Sprintf(format, args...)
	
	var callerInfo string
	if l.showCaller {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			file = filepath.Base(file)
			callerInfo = fmt.Sprintf(" [%s:%d]", file, line)
		}
	}
	
	logEntry := fmt.Sprintf("%s [%s] %s%s: %s\n", 
		timestamp, 
		levelNames[level], 
		l.prefix,
		callerInfo,
		message)
	
	fmt.Fprintf(l.output, logEntry)
	
	// If FATAL, exit the program
	if level == FATAL {
		os.Exit(1)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
}

// WithFields returns a formatted string for structured logging
func WithFields(fields map[string]interface{}) string {
	parts := make([]string, 0, len(fields))
	for k, v := range fields {
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(parts, " ")
}

// Global logger functions

// SetDefaultLevel sets the level of the default logger
func SetDefaultLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// Debug logs a debug message to the default logger
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info logs an info message to the default logger
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warn logs a warning message to the default logger
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Error logs an error message to the default logger
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// Fatal logs a fatal message to the default logger and exits the program
func Fatal(format string, args ...interface{}) {
	defaultLogger.Fatal(format, args...)
}
