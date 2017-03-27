package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger defines the public interface of a logger as defined by this package
type Logger interface {
	Infoln(argv ...interface{})
	Infof(format string, argv ...interface{})
	Errorln(argv ...interface{})
	Errorf(format string, argv ...interface{})
}

// New creates a new logger,
// using a file if it is defined, and otherwise the STDERR
func New(path string, verbose bool) Logger {
	w := os.Stderr
	if path != "" {
		lf, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
		// err is not logged, as there is no logger yet
		if err == nil {
			w = lf
		}
	}

	l := log.New(w, "go-gitignore",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.LUTC)
	return &logger{logger: l, verbose: verbose}
}

type logger struct {
	logger  *log.Logger
	verbose bool
}

// Infoln implements Logger.Infoln
func (l *logger) Infoln(argv ...interface{}) {
	if !l.verbose {
		return
	}

	l.logger.Println("[INFO] " + fmt.Sprint(argv...))
}

// Infof implements Logger.Infof
func (l *logger) Infof(format string, argv ...interface{}) {
	if !l.verbose {
		return
	}

	l.logger.Printf("[INFO] "+format, argv...)
}

// Errorln implements Logger.Errorln
func (l *logger) Errorln(argv ...interface{}) {
	l.logger.Println("[ERROR] " + fmt.Sprint(argv...))
}

// Errorf implements Logger.Errorf
func (l *logger) Errorf(format string, argv ...interface{}) {
	l.logger.Printf("[ERROR] "+format, argv...)
}
