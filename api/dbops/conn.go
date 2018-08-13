/**
 * @Author codeAC
 * @Time: 2018/8/12 18:54
 * @Description
 */
package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123@tcp(localhost:3306)/videoserver?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
