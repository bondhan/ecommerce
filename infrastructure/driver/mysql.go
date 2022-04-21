package driver

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bondhan/ecommerce/infrastructure/config"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

// GormLogger is a custom Logger for Gorm, making it use logrus.
type GormLogger struct {
	Logger                *logrus.Logger
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func (l *GormLogger) LogMode(level log.LogLevel) log.Interface {
	if level <= log.Warn {
		l.Logger.SetLevel(logrus.WarnLevel)
	} else {
		l.Logger.SetLevel(logrus.TraceLevel)
	}
	return l
}
func (l *GormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.WithContext(ctx).Info(msg, args)
}
func (l *GormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.WithContext(ctx).Warn(msg, args)
}
func (l *GormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.WithContext(ctx).Error(msg, args)
}
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sqlVar, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		l.Logger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sqlVar, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.Logger.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sqlVar, elapsed)
		return
	}

	l.Logger.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sqlVar, elapsed)
}

// NewDBInstance ...
func NewDBInstance(logger *logrus.Logger, confDsnMaster string, poolType config.ConnPoolType) (*gorm.DB, *sql.DB) {

	//lvl := logger.GetLevel()
	//var level log.LogLevel
	//switch lvl {
	//case logrus.ErrorLevel:
	//	level = log.Error
	//case logrus.WarnLevel:
	//	level = log.Warn
	//case logrus.InfoLevel:
	//	level = log.Info
	//default:
	//	level = log.Info
	//}
	//
	//l := &GormLogger{Logger: logger}

	db, err := gorm.Open(mysql.Open(confDsnMaster), &gorm.Config{
		//Logger: l.LogMode(level),
	})
	if err != nil {
		logger.Fatalf("fail open database err:%s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("database err:%s", err)
	}

	if poolType.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(poolType.MaxOpenConn)
	}

	if poolType.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(poolType.MaxIdleConn)
	}

	if poolType.MaxLifeTime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(poolType.MaxLifeTime) * time.Minute)
	}

	return db, sqlDB
}
