package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Connection *sql.DB
var err error

func Connect(DBTYPE, DBHOST, DBUSER, DBPASS, DBNAME, DBPORT string) (*sql.DB, error) {
	fmt.Println(DBTYPE, DBHOST, DBUSER, DBPASS, DBNAME, DBPORT)
	stringConnection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DBUSER, DBPASS, DBHOST, DBPORT, DBNAME,
	)

	Connection, err = sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}
	if err = Connection.Ping(); err != nil {
		Connection.Close()
		return nil, err
	}
	return Connection, nil
}
