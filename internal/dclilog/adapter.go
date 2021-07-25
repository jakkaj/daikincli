package dclilog

//Version is the dcli overall build version
var Version = "edge"

// Logger struct wrapping around an Adapter.
type Logger struct {
	adapter Adapter
}

// Adapter interface for different log levels.
type Adapter interface {
	Printf(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Error(...interface{})
}

// SetLogger changes the adapter used by the logger.
func (l *Logger) SetLogger(a Adapter) {
	l.adapter = a
}

// Printf is used for printing out messages for user.
func (l *Logger) Printf(fmt string, args ...interface{}) {
	l.adapter.Printf(fmt, args...)
}

// Debugf is used for logging debug level logs.
func (l *Logger) Debugf(fmt string, args ...interface{}) {
	l.adapter.Debugf(fmt, args...)
}

// Infof is used for logging info level logs.
func (l *Logger) Infof(fmt string, args ...interface{}) {
	l.adapter.Infof(fmt, args...)
}

// Warnf is used for logging warning level logs.
func (l *Logger) Warnf(fmt string, args ...interface{}) {
	l.adapter.Warnf(fmt, args...)
}

// Errorf is used for logging error level logs.
func (l *Logger) Errorf(fmt string, args ...interface{}) {
	l.adapter.Errorf(fmt, args...)
}

// Error is used for logging error level logs.
func (l *Logger) Error(args ...interface{}) {
	l.adapter.Error(args...)
}

// NewLogger creates a new instance of Logger with passed in adapter.
func NewLogger(a Adapter) Logger {
	return Logger{adapter: a}
}
