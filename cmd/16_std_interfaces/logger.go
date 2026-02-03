package stdinterfaces

import (
	"fmt"
	"io"
	"time"
)

type FormatError struct {
	Level string
}

type LogEntry struct {
	Level     string
	Message   string
	Timestamp time.Time
}

type LogWriter struct {
	buffer []byte
}

type LogReader struct {
	data   []byte
	offset int
}

func (l *LogEntry) String() string {
	return fmt.Sprintf("[%s] %v - %s", l.Level, l.Timestamp.Format("2006-01-02 15:04:05"), l.Message)
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	w.buffer = append(w.buffer, p...)
	return len(p), nil
}

func (w *LogWriter) String() string {
	return string(w.buffer)
}

func (r *LogReader) Read(p []byte) (int, error) {
	if r.offset >= len(r.data) {
		return 0, io.EOF
	}

	n := copy(p, r.data[r.offset:])
	r.offset += n

	return n, nil
}

func (e *FormatError) Error() string {
	return fmt.Sprintf("Invalid Log Level: %s", e.Level)
}

func NewLogReader(data []byte) *LogReader {
	return &LogReader{data: data}
}

func NewLogEntry(level, message string) (LogEntry, error) {
	switch level {
	case "INFO", "WARN", "ERROR":
		return LogEntry{
			Level:     level,
			Message:   message,
			Timestamp: time.Now(),
		}, nil
	default:
		return LogEntry{}, &FormatError{Level: level}
	}
}
