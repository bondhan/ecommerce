package driver

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
	"time"
)

type DefaultFieldHook struct {
	fields map[string]interface{}
}

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	for i, v := range h.fields {
		e.Data[i] = v
	}
	return nil
}

// NewLogInstance ...
func NewLogInstance(isProd bool, fields map[string]interface{}) *logrus.Logger {
	var level logrus.Level

	logger := logrus.New()

	//if it is production will output warn and error level
	if isProd {
		level = logrus.WarnLevel
	} else {
		level = logrus.TraceLevel
	}

	logger.SetLevel(level)
	logger.SetOutput(colorable.NewColorableStdout())
	if isProd {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
			PrettyPrint:     true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			//PrettyPrint:     true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
	}

	logger.AddHook(&DefaultFieldHook{fields})

	return logger
}
