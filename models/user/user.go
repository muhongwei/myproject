package user

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/golang/glog"
	//"github.com/astaxie/beedb"
	_ "github.com/Go-SQL-Driver/MySQL"
)

//User User
type User struct {
	UserId           int    `json:"userid,omitempty"`
	UserName         string `json:"username,omitempty"`
	UserPassword     string `json:"userpassword,omitempty"`
	UserIntroduction string `json:"userintroduction,omitempty"`
}

//初始化一个数据库连接
func initMysql() *sql.DB {
	//打开数据库连接Open(驱动名,连接字符串)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/myapp?charset=utf8")
	if err != nil {
		glog.Fatalln(err)
	}
	return db
}

//FindUser FindUser
func FindUser(user User) *User {
	fmt.Println("finduser")
	db := initMysql()
	defer db.Close()
	row, err := db.Query("select * from `myapp`.`user`")
	if err != nil {
		glog.Fatalln(err)
	}
	var user1 User
	for row.Next() {
		row.Scan(&user1.UserId, &user1.UserName, &user1.UserPassword, &user1.UserIntroduction)
		// log.Println("id:", id, ",name:", name, "password:", password,"introduction:",introduction)
		if (user1.UserName == user.UserName) && (user1.UserPassword == user.UserPassword) {
			return &user1
		}
	}
	return nil

}

//DeleteUser DeleteUser
func DeleteUser(user User) error {
	db := initMysql()
	defer db.Close()
	sql := "delete  from `myapp`.`user` where userId = " + strconv.Itoa(user.UserId)
	glog.Infoln(sql)
	_, err := db.Query(sql)
	if err != nil {
		glog.Fatalln(err)
	}
	return err
}

//SaveUser SaveUser
func SaveUser(user User) error {
	db := initMysql()
	defer db.Close()
	fmt.Println(user)
	_, err := db.Exec("insert into `myapp`.`user`(userName,userPassword,userIntroduction) values(?,?,?)", user.UserName, user.UserPassword, user.UserIntroduction)
	return err
}

//ValidateUser ValidateUser
func ValidateUser(user User) error {
	db := initMysql()
	defer db.Close()
	row, err := db.Query("select userName,userPassword from `myapp`.`user`")
	if err != nil {
		glog.Fatalln(err)
	}
	var name string = ""
	var password string = ""
	for row.Next() {
		row.Scan(&name, &password)
		// log.Println("id:", id, ",name:", name, "password:", password,"introduction:",introduction)
		if (name == user.UserName) && (password == user.UserPassword) {
			return nil
		}
	}
	return errors.New("用户名或密码错误！")

}

//GetAllUser GetAllUser
func GetAllUser() []User {
	users := []User{}
	db := initMysql()
	defer db.Close()
	row, err := db.Query("select userId,userName,userPassword,userIntroduction from `myapp`.`user`")
	if err != nil {
		glog.Fatalln(err)
	}
	var id int
	var name string = ""
	var pwd string = ""
	var intro string = ""
	for row.Next() {
		row.Scan(&id, &name, &pwd, &intro)
		// log.Println("id:", id, ",name:", name, "password:", password,"introduction:",introduction)
		var user User
		user.UserId = id
		user.UserName = name
		user.UserPassword = pwd
		user.UserIntroduction = intro
		users = append(users, user)

	}
	return users
}
