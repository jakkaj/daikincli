package dclilog

import (
	f "fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var (
	lock = sync.RWMutex{}
	log  *Logger
	once sync.Once
)

// LogrusAdapter is used as a passable adapter used to log.
type LogAdapter struct{}

// Printf prints out messages for user, appending line break at the end
func (l LogAdapter) Printf(fmt string, args ...interface{}) {
	lock.Lock()
	f.Printf(f.Sprintln(fmt), args...)
	f.Println("")
	lock.Unlock()
}

// Debugf is used for logging debug level logs.
func (l LogAdapter) Debugf(fmt string, args ...interface{}) {
	color.Set(color.FgMagenta)
	lock.Lock()
	f.Printf(fmt, args...)
	f.Println("")
	lock.Unlock()
	color.Unset()
}

// Warnf is used for logging warning level logs.
func (l LogAdapter) Warnf(fmt string, args ...interface{}) {
	color.Set(color.FgYellow)
	lock.Lock()
	f.Printf(fmt, args...)
	f.Println("")
	lock.Unlock()
	color.Unset()
}

// Infof is used for logging info level logs.
func (l LogAdapter) Infof(fmt string, args ...interface{}) {
	lock.Lock()
	f.Printf(fmt, args...)
	f.Println("")
	lock.Unlock()
	color.Unset()
}

// Errorf is used for logging error level logs.
func (l LogAdapter) Errorf(fmt string, args ...interface{}) {
	color.Set(color.FgRed)
	lock.Lock()
	logrus.Errorf(fmt, args...)
	lock.Unlock()
	color.Unset()
}

// Error is used for logging error level logs.
func (l LogAdapter) Error(args ...interface{}) {
	color.Set(color.FgRed)
	lock.Lock()
	logrus.Error(args...)
	lock.Unlock()
	color.Unset()
}

// GetInstance returns an instance of adapter.Logger
func GetInstance() *Logger {
	once.Do(func() {
		logger := NewLogger(LogAdapter{})
		log = &logger

	})
	return log
}
