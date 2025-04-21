package initializer

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, msg string) {
	if err != nil {
		global.Logger.Error(msg, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DatabaseName)
	db, err := sql.Open(global.Config.Server.DriverName, s)
	checkErrorPanic(err, "InitMysql error")
	global.Logger.Info("InitMysql success")
	global.Mysql = db
}
