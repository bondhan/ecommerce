package main

import (
	"database/sql"
	"github.com/bondhan/ecommerce/infrastructure/config"
	"github.com/bondhan/ecommerce/infrastructure/driver"
	"github.com/bondhan/ecommerce/interfaces/http"
	"github.com/bondhan/ecommerce/migrations"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	// instantiate log
	isProd, m := config.NewLogConf(os.Getenv("ENV"), os.Getenv("APP_NAME"))
	logger := driver.NewLogInstance(isProd, m)

	// create database if not exist
	migrations.CreateDB(logger, os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DBNAME"))

	// instantiate database
	dsn := config.NewDsnMYSQLDBConf(os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DBNAME"))
	poolType := config.NewConnPoolConf(logger, os.Getenv("DB_MAX_OPEN_CONN"), os.Getenv("DB_MAX_IDLE_CONN"),
		os.Getenv("DB_MAX_LIFE_TIME_CONN_MINUTE"))
	logger.Debugf(dsn)
	gormDBConn, sqlDBConn := driver.NewDBInstance(logger, dsn, poolType)

	defer func(l *logrus.Logger, sqlDb *sql.DB) {
		err := sqlDb.Close()
		if err != nil {
			l.Errorf("Error closing sqlDbData: %s", err)
		} else {
			l.Warn("DB closed successfully")
		}
	}(logger, sqlDBConn)

	// migration up
	if err := migrations.MigrateMysqlUp(logger, sqlDBConn); err != nil {
		logger.Errorf("migrate err: %s", err)
		return
	}

	service := handler.NewHandler(logger, os.Getenv("JWT_KEY"), gormDBConn)
	router := handler.NewRouter(logger, service, os.Getenv("JWT_KEY"))

	driver.RunHttpServer(logger, os.Getenv("PORT"), router)
}
