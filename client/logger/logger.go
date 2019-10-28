package logger

import (
	"fmt"
	"github.com/shyamz-22/syslog-client/exceptions"
	"io/ioutil"
	"log"
	"log/syslog"
	"os"
)

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
	Off
)

type Logger struct {
	debugLogger          *log.Logger
	infoLogger           *log.Logger
	warnLogger           *log.Logger
	errorLogger          *log.Logger
	fatalLogger          *log.Logger
	sysLogger            *syslog.Writer
	mandatoryErrorLogger *log.Logger
	level                Level
}

func NewLogger(l Level) *Logger {
	flag := log.LstdFlags | log.Lmicroseconds | log.Lshortfile
	debugWriter := ioutil.Discard
	infoWriter := ioutil.Discard
	warnWriter := ioutil.Discard
	errorWriter := ioutil.Discard
	fatalWriter := ioutil.Discard

	switch l {
	case Debug:
		debugWriter = os.Stdout
		infoWriter = os.Stdout
		warnWriter = os.Stdout
		errorWriter = os.Stderr
		fatalWriter = os.Stderr
	case Info:
		infoWriter = os.Stdout
		warnWriter = os.Stdout
		errorWriter = os.Stderr
		fatalWriter = os.Stderr
	case Warn:
		warnWriter = os.Stdout
		errorWriter = os.Stderr
		fatalWriter = os.Stderr
	case Error:
		errorWriter = os.Stderr
		fatalWriter = os.Stderr
	case Fatal:
		fatalWriter = os.Stderr
	}

	return &Logger{
		debugLogger:          log.New(debugWriter, "Debug: ", flag),
		infoLogger:           log.New(infoWriter, "Info: ", flag),
		warnLogger:           log.New(warnWriter, "Warn: ", flag),
		errorLogger:          log.New(errorWriter, "Error: ", flag),
		fatalLogger:          log.New(fatalWriter, "Fatal: ", flag),
		mandatoryErrorLogger: log.New(os.Stderr, "Error: ", flag),
		level:                l,
	}

}

func (l *Logger) SetSysLogger(sl *syslog.Writer) {
	l.sysLogger = sl
}

func (l *Logger) Debug(v ...interface{}) {
	message := fmt.Sprint(v)
	l.debugLogger.Println(message)

	if Debug == l.level {
		exceptions.LogAndIgnore(l.mandatoryErrorLogger, l.sysLogger.Debug(message))
	}

}

func (l *Logger) Debugf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v)
	l.Debug(message)
}

func (l *Logger) Info(v ...interface{}) {
	message := fmt.Sprint(v)
	l.infoLogger.Println(message)

	if l.level < Warn {
		exceptions.LogAndIgnore(l.mandatoryErrorLogger, l.sysLogger.Info(message))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v)
	l.Info(message)
}

func (l *Logger) Warn(v ...interface{}) {
	message := fmt.Sprint(v)
	l.warnLogger.Println(message)

	if l.level < Error {
		exceptions.LogAndIgnore(l.mandatoryErrorLogger, l.sysLogger.Warning(message))
	}

}

func (l *Logger) Warnf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v)
	l.Warn(message)
}

func (l *Logger) Error(v ...interface{}) {
	message := fmt.Sprint(v)
	l.errorLogger.Println(message)

	if l.level < Fatal {
		exceptions.LogAndIgnore(l.mandatoryErrorLogger, l.sysLogger.Err(message))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v)
	l.Warn(message)
}

func (l *Logger) Fatal(v ...interface{}) {
	message := fmt.Sprint(v)
	l.fatalLogger.Println(message)

	if l.level < Off {
		exceptions.LogAndIgnore(l.mandatoryErrorLogger, l.sysLogger.Emerg(message))
	}
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v)
	l.Fatal(message)
}
