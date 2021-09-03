package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Backup Process")

}

//db related
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "atenobserve"
	dbname   = "testDB"
)

type userinfo struct {
	uid        int
	username   string
	departname string
	Created    string `sql:"type:timestamp"`
}

type appContext struct {
	db *sql.DB
}

func connectDB() (c *appContext, errorMessage string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	   "password=%s dbname=%s sslmode=disable",
	   host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	   fmt.Println("连接数据库错误："+err.Error())
	}else {
	   fmt.Println("Successfully connected!")
	}
 
	err = db.Ping()
	if err != nil {
	   fmt.Println("DBPing错误："+err.Error())
	}else
	{
	   fmt.Println("DBPingSuccess")
 	}
	return &appContext{db}, ""
 

func lockDb() {
	return
}

func unlockDb() {
	return
}

func uploadS3() {
	return
}
