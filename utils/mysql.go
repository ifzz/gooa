package utils

import (
	"database/sql"
	"github.com/go-gorp/gorp/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
	"log"
	"os"
)

func InitDingDb() *gorp.DbMap {
	db, err := sql.Open("mysql", DingConnection)
	CheckErr(err, "sql.Open failed")

	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	if TraceOn {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "[SQL]:", log.Lmicroseconds))
	}

	return dbMap
}

func InitIisyDb() *gorp.DbMap {
	db, err := sql.Open("mysql", IissyConnection)
	CheckErr(err, "sql.Open failed")

	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	if TraceOn {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "[SQL]:", log.Lmicroseconds))
	}

	return dbMap
}

func BuildSqlArgs(args ...interface{}) ([]interface{}, error) {
	newArgs := make([]interface{}, 0)
	addEleFun := func(ele interface{}) {
		newArgs = append(newArgs, ele)
		return
	}
	for _, arg := range args {
		switch v := arg.(type) {
		case string, int, int32, int64, bool, *string, *int, *int32, *int64:
			addEleFun(v)
		case []string:
			for _, e := range v {
				addEleFun(e)
			}
		case []int:
			for _, e := range v {
				addEleFun(e)
			}
		case []int32:
			for _, e := range v {
				addEleFun(e)
			}
		case []int64:
			for _, e := range v {
				addEleFun(e)
			}
		case []*string:
			for _, e := range v {
				addEleFun(e)
			}
		default:
			return nil, errors.New("miss type")
		}
	}
	return newArgs, nil
}
