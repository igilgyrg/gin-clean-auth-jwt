package logging

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var logger *logrus.Entry

func newLogger() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s %d", filename, frame.Line), fmt.Sprintf("%s", frame.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	l.SetOutput(os.Stdout)

	for _, level := range logrus.AllLevels {
		l.SetLevel(level)
	}

	logger = logrus.NewEntry(l)
}

func Logger() *logrus.Entry {
	once.Do(newLogger)
	return logger
}