package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-gorp/gorp"

	. "github.com/apdaza/oasRuler/models"
	. "github.com/apdaza/oasRuler/utils"
)

const (
	dbHost = "tcp(127.0.0.1:3306)"
	dbName = "rules"
	dbUser = "ruleuser"
	dbPass = "ru13_u53r"
)
/*InitDb es usuada para iniciar la conexion a base de datos*/
func InitDb() *gorp.DbMap {
	dsn := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?charset=utf8"

	db, err := sql.Open("mysql", dsn)

	CheckErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Domain{}, "domain").SetKeys(true, "Id")
	dbmap.AddTableWithName(Rule{}, "rule").SetKeys(true, "Id")
	dbmap.AddTableWithName(Component{}, "component").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}
