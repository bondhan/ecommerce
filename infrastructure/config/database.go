package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ConnPoolType struct {
	MaxOpenConn int
	MaxIdleConn int
	MaxLifeTime int64
}

// NewDsnMYSQLDBConf will return data source naming for MYSQL
func NewDsnMYSQLDBConf(dbHost, dbPort, dbUser, dbPassword, dbName string) string {

	dsnDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", dbUser,
		dbPassword, dbHost, dbPort, dbName)

	return dsnDB
}

// NewConnPoolConf will return the consolidated connection pool
func NewConnPoolConf(logger *logrus.Logger, maxOpen string, maxIdle string, maxLifeTimeMin string) ConnPoolType {
	maxOpenConn, err := strconv.Atoi(maxOpen)
	if err != nil {
		logger.Fatalf("maxOpen: %s err: %s", maxOpen, err)
	}

	maxIdleConn, err := strconv.Atoi(maxIdle)
	if err != nil {
		logger.Fatalf("maxIdle: %s err: %s", maxIdle, err)
	}

	maxLifetime, err := strconv.Atoi(maxLifeTimeMin)
	if err != nil {
		logger.Fatalf("maxLifetime: %s err: %s", maxLifeTimeMin, err)
	}

	return ConnPoolType{
		MaxOpenConn: maxOpenConn,
		MaxIdleConn: maxIdleConn,
		MaxLifeTime: int64(maxLifetime),
	}
}
