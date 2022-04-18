package query

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type cashierQ struct {
	logger *logrus.Logger
	gormDB *gorm.DB
}

func NewCashierQ(logger *logrus.Logger, gDB *gorm.DB) ICashierQ {
	return &cashierQ{
		logger: logger,
		gormDB: gDB,
	}
}
