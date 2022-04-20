package migrations

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

const Mysql = "mysql"

type migrateLog struct {
	isVerbose bool
	logger    *logrus.Logger
}

func (l *migrateLog) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *migrateLog) Verbose() bool {
	return l.isVerbose
}

func MigrateMysqlUp(logger *logrus.Logger, db *sql.DB) error {

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logger.Errorf("error instantiate driver err:%s", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/scripts",
		Mysql,
		driver,
	)

	m.Log = &migrateLog{logger: logger, isVerbose: false}

	if err != nil {
		logger.Errorf("error instantiate migrate err: %s", err)
		return err
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			logger.Warnf("migration: %s", err)
			return nil
		}
		logger.Errorf("error migrating up err: %s", err)
		return err
	}

	return nil
}

func CreateDB(logger *logrus.Logger, host, port, user, password, dbname string) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)

	db, err := sql.Open(Mysql, dsn)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		logger.Fatal(err)
	}

	_, err = db.Exec("USE " + dbname)
	if err != nil {
		logger.Fatal(err)
	}

}
