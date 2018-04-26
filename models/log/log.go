package log

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	//"github.com/astaxie/beedb"
	_ "github.com/Go-SQL-Driver/MySQL"
)

//Log Log
type Log struct {
	ID       int    `json:"id,omitempty"`
	Time     string `json:"time"`
	UserName string `json:"username"`
	API      string `json:"api"`
	Method   string `json:"method"`
}
type LOGC struct {
	beego.Controller
}

//初始化一个数据库连接
func initMysql() *sql.DB {
	//打开数据库连接Open(驱动名,连接字符串)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/myapp?charset=utf8&loc=Local")
	if err != nil {
		glog.Fatalln(err)
	}
	return db
}

//InsertLog InsertLog
func InsertLog(log Log) error {
	db := initMysql()
	defer db.Close()
	local2, err2 := time.LoadLocation("Asia/Chongqing") //服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	glog.Infoln("InsertLog")
	_, err := db.Exec("insert into `myapp`.`log`(time,user,api,method) values(?,?,?,?)", time.Now().In(local2), log.UserName, log.API, log.Method)
	return err
}

//ListLog ListLog
func ListLog() []Log {
	db := initMysql()
	defer db.Close()
	row, err := db.Query("select time,user,api,method from `myapp`.`log` order by time desc limit 100")
	if err != nil {
		glog.Fatalln(err)
		return nil
	}
	var logs []Log
	var log Log

	for row.Next() {
		row.Scan(&log.Time, &log.UserName, &log.API, &log.Method)
		logs = append(logs, log)

	}
	return logs
}

// func Loglog(){
// 	var vlog Log
// 	var c = LOGC{}
// 	sess := c.StartSession()
// 	username := sess.Get("username")
// 	if username == nil {
// 		c.CustomAbort(400, "login first")
// 		return
// 	} else {
// 		vlog.UserName = interface{}(username).(string)
// 		vlog.API = "/user/replicationcontroller/list"
// 		vlog.Method = "get"
// 		InsertLog(vlog)

// 	}
// }
