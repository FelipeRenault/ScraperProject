package db

import (
	"database/sql"
	"time"

	"github.com/go-gorp/gorp"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	dsn         = "MainUser:MainPassword@tcp(pelandochallenge-mysql-1:3306)/pelando?charset=utf8&parseTime=true&loc=America%2FSao_Paulo"
	maxLifetime = 10 * time.Second
)

func Connect() (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	db.SetConnMaxLifetime(maxLifetime)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return dbmap, nil
}
