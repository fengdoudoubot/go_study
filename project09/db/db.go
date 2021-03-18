package db

import (
	"database/sql"
	_"fmt"
	_"github.com/go-sql-driver/mysql"
)
var(
	Db *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@/test")
	if err != nil {
		panic(err.Error())
	}

}
