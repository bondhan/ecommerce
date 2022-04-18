package main

import (
	"database/sql"
	"github.com/bondhan/ecommerce/infrastructure/config"
	"github.com/bondhan/ecommerce/infrastructure/driver"
	"github.com/bondhan/ecommerce/interfaces/http"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	// instantiate log
	isProd, m := config.NewLogConf(os.Getenv("ENV"), os.Getenv("APP_NAME"))
	logger := driver.NewLogInstance(isProd, m)

	// instantiate database
	dsn := config.NewDsnMYSQLDBConf(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	poolType := config.NewConnPoolConf(logger, os.Getenv("DB_MAX_OPEN_CONN"), os.Getenv("DB_MAX_IDLE_CONN"),
		os.Getenv("DB_MAX_LIFE_TIME_CONN_MINUTE"))
	gormDBConn, slqDBConn := driver.NewDBInstance(logger, dsn, poolType)

	defer func(l *logrus.Logger, sqlDb *sql.DB) {
		err := sqlDb.Close()
		if err != nil {
			l.Errorf("Error closing sqlDbData: %s", err)
		}
	}(logger, slqDBConn)

	_ = gormDBConn

	service := handler.NewHandler(logger, gormDBConn)
	handler := handler.NewRouter(service)

	driver.RunHttpServer(logger, os.Getenv("PORT"), handler)
}
