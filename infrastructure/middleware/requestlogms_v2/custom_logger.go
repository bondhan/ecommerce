package requestlogmsv2

import (
	"bytes"
	"encoding/json"
	requestlogmsv2_domain "github.com/bondhan/ecommerce/infrastructure/middleware/requestlogms_v2/domain"
	"github.com/bondhan/ecommerce/infrastructure/utilities"
	"gorm.io/gorm"

	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/guregu/null"
	"github.com/sirupsen/logrus"
)

const BufferSizeInBytes = 1 * 1024

var (
	dbLog *gorm.DB
	mutex sync.Mutex
)

// CustomLoggerV2 ....
func CustomLoggerV2(logger *logrus.Logger, dbConn *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			entry := &StructuredLoggerEntry{Logger: logger, Data: make(map[string]interface{})}
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			buf := utilities.NewLimitBuffer(BufferSizeInBytes)
			ww.Tee(buf)

			//get and record the raw and encrypted body
			var body null.String

			bdy, err := ioutil.ReadAll(io.LimitReader(r.Body, BufferSizeInBytes))
			if err != nil {
				logger.Error("Fail Reading Request Body:", err)
			}

			if len(bdy) > 0 {
				body = null.StringFrom(string(bdy))
			}

			// And now set a new body, which will simulate the same data we read:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bdy))

			//get and record the requestID
			reqID := middleware.GetReqID(r.Context())

			//get and record the raw header as json
			header, err := json.Marshal(r.Header)
			if err != nil {
				logger.Error("Fail Marshall Header Data:", err)
			}

			defer func(t time.Time, ent *StructuredLoggerEntry) {
				data := ent.Data
				trace := ""
				if data["trace"] != nil {
					trace = data["trace"].(string)
				}

				//get and record the raw and encrypted body
				rBody, err := ioutil.ReadAll(buf)
				if err != nil {
					logger.Error("Fail Reading Response Body:", err)
				}
				respBody := null.StringFrom(string(rBody))

				//assign one time if and only if dbConn was not nil
				if dbConn != nil && dbLog == nil {
					mutex.Lock()
					dbLog = dbConn
					mutex.Unlock()
				}

				if dbLog != nil {
					// compose the object for logging
					log := requestlogmsv2_domain.LogRequest{
						IPAddress:    null.StringFrom(r.RemoteAddr),
						Header:       null.StringFrom(string(header)),
						Host:         null.StringFrom(r.Host),
						URL:          null.StringFrom(r.URL.Path),
						URI:          null.StringFrom(r.URL.RequestURI()),
						HTTPMethod:   null.StringFrom(r.Method),
						HTTPRespCode: null.IntFrom(int64(ww.Status())),
						Message:      null.StringFrom(trace),
						UserAgent:    null.StringFrom(r.UserAgent()),
						RequestID:    null.StringFrom(reqID),
						RequestTime:  null.TimeFrom(t),
						ResponseTime: null.TimeFrom(time.Now().UTC()),
						LapsedTimeMs: null.IntFrom(time.Since(t).Milliseconds()),
						RequestBody:  body,
						ResponseBody: respBody,
					}

					err = dbLog.Create(&log).Error
					if err != nil {
						e := make(map[string]interface{})
						e["message"] = "Fail Insert to Log DB"
						e["error"] = err

						logger.Error(e)
					}
				}

				/***** Logrus ***/
				logFields := logrus.Fields{}

				logFields["type"] = "request"
				logFields["requestId"] = reqID

				data["ip"] = null.StringFrom(r.RemoteAddr)
				data["userAgent"] = null.StringFrom(r.UserAgent())
				data["method"] = null.StringFrom(r.Method)
				data["url"] = null.StringFrom(r.URL.Path)
				data["headers"] = null.StringFrom(string(header))
				data["params"] = ""
				data["query"] = null.StringFrom(r.URL.RawQuery)
				data["body"] = body
				data["response"] = respBody
				data["status"] = null.IntFrom(int64(ww.Status()))

				logFields["data"] = data

				if ent.Level == logrus.ErrorLevel || ent.Level == logrus.WarnLevel {
					ent.Logger.WithFields(logFields).Log(ent.Level, "request log")
				} else {
					ent.Logger.WithFields(logFields).Info("request log")
				}

			}(time.Now().UTC(), entry)

			next.ServeHTTP(ww, middleware.WithLogEntry(r, entry))
		})
	}
}

// CustomLoggerV2PerHandler ....
func CustomLoggerV2PerHandler(logger *logrus.Logger, dbConn *gorm.DB, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entry := &StructuredLoggerEntry{Logger: logger, Data: make(map[string]interface{})}
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		buf := utilities.NewLimitBuffer(BufferSizeInBytes)
		ww.Tee(buf)

		//get and record the raw and encrypted body
		var body null.String

		bdy, err := ioutil.ReadAll(io.LimitReader(r.Body, BufferSizeInBytes))
		if err != nil {
			logger.Error("Fail Reading Request Body:", err)
		}

		if len(bdy) > 0 {
			body = null.StringFrom(string(bdy))
		}

		// And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bdy))

		//get and record the requestID
		reqID := middleware.GetReqID(r.Context())

		//get and record the raw header as json
		header, err := json.Marshal(r.Header)
		if err != nil {
			logger.Error("Fail Marshall Header Data:", err)
		}

		defer func(t time.Time, ent *StructuredLoggerEntry) {
			data := ent.Data
			trace := ""
			if data["trace"] != nil {
				trace = data["trace"].(string)
			}

			//get and record the raw and encrypted body
			rBody, err := ioutil.ReadAll(buf)
			if err != nil {
				logger.Error("Fail Reading Response Body:", err)
			}
			respBody := null.StringFrom(string(rBody))

			//assign one time if and only if dbConn was not nil
			if dbConn != nil && dbLog == nil {
				mutex.Lock()
				dbLog = dbConn
				mutex.Unlock()
			}

			if dbLog != nil {
				// compose the object for logging
				log := requestlogmsv2_domain.LogRequest{
					IPAddress:    null.StringFrom(r.RemoteAddr),
					Header:       null.StringFrom(string(header)),
					Host:         null.StringFrom(r.Host),
					URL:          null.StringFrom(r.URL.Path),
					URI:          null.StringFrom(r.URL.RequestURI()),
					HTTPMethod:   null.StringFrom(r.Method),
					HTTPRespCode: null.IntFrom(int64(ww.Status())),
					Message:      null.StringFrom(trace),
					UserAgent:    null.StringFrom(r.UserAgent()),
					RequestID:    null.StringFrom(reqID),
					RequestTime:  null.TimeFrom(t),
					ResponseTime: null.TimeFrom(time.Now().UTC()),
					LapsedTimeMs: null.IntFrom(time.Since(t).Milliseconds()),
					RequestBody:  body,
					ResponseBody: respBody,
				}

				err = dbLog.Create(&log).Error
				if err != nil {
					e := make(map[string]interface{})
					e["message"] = "Fail Insert to Log DB"
					e["error"] = err

					logger.Error(e)
				}
			}

			/***** Logrus ***/
			logFields := logrus.Fields{}

			logFields["type"] = "request"
			logFields["requestId"] = reqID

			data["ip"] = null.StringFrom(r.RemoteAddr)
			data["userAgent"] = null.StringFrom(r.UserAgent())
			data["method"] = null.StringFrom(r.Method)
			data["url"] = null.StringFrom(r.URL.Path)
			data["headers"] = null.StringFrom(string(header))
			data["params"] = ""
			data["query"] = null.StringFrom(r.URL.RawQuery)
			data["body"] = body
			data["response"] = respBody
			data["status"] = null.IntFrom(int64(ww.Status()))

			logFields["data"] = data

			if ent.Level == logrus.ErrorLevel || ent.Level == logrus.WarnLevel {
				ent.Logger.WithFields(logFields).Log(ent.Level, "request log")
			} else {
				ent.Logger.WithFields(logFields).Info("request log")
			}

		}(time.Now().UTC(), entry)

		next.ServeHTTP(ww, middleware.WithLogEntry(r, entry))
	})
}
