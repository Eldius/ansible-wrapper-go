package logger

import (
	"log"
)

/*
LogWriter wrapper for log to use at os.exec.Cmd
*/
type LogWriter struct {
	logger *log.Logger
}

/*
NewLogWriter creates a new LogWriter
*/
func NewLogWriter(l *log.Logger) *LogWriter {
	return &LogWriter{
		logger: l,
	}
}

/*
DefaultLogWriter creates a new LogWriter
*/
func DefaultLogWriter() *LogWriter {
	return &LogWriter{
		logger: DefaultLogger(),
	}
}

/*
DefaultLogger returns a default logger
*/
func DefaultLogger() *log.Logger {
	return log.New(
		log.Writer(),
		log.Prefix()+"> ",
		log.Flags(),
	)
}

/*
Write write logs to log (doh)
*/
func (lw LogWriter) Write(p []byte) (n int, err error) {
	lw.logger.Print(string(p))
	return len(p), nil
}
