package database

import (
	"database/sql"
	"fmt"

	"github.com/celmysql-api/utils"
	_ "github.com/go-sql-driver/mysql"
)

// func GetConnection(config utils.Config) (db *sql.DB) {
// 	// dbDriver := "mysql"
// 	// dbUser := config.Dbuser
// 	// dbPass := config.Database
// 	// dbName := config.Dbpassword
// 	datasource := "root:@tcp(localhost:3306)/bk3devdb?parseTime=true"
// 	db, err := sql.Open("mysql", datasource)
// 	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// func GetConnection(config utils.Config) (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := "Serverh5n&*#"

// 	// dbName := "bk3db"
// 	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
// 	dbName := "(103.157.96.47:3306)/bk3db"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp"+dbName+"?parseTime=true")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

func GetConnection(config utils.Config) (db *sql.DB) {
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.MysqlDBUser, config.MysqlDBPass, config.MysqlDBHost, config.MysqlDBPort, config.MysqlDBName)
	dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "Serverh5n&*#"

	// dbName := "bk3db"
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	// dbName := "(localhost:3306)/celestialmysqldb"
	db, err := sql.Open(dbDriver, mysqlDSN)
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

// func TestOpenConnection(t *testing.T) {
// 	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := "Serverh5n&*#"
// 	dbName := "celestialmysqldb"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// gunakan DB
// }
