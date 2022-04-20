package requestlogmsv2

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
	Data   map[string]interface{}
	Level  logrus.Level
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {

}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {

}
