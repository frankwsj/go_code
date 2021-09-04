package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Backup Process")
	c, err := connectDB()
	fmt.Println(c)
	defer c.db.Close()

	if err != "" {
		print(err)
	}
	//c.Create()
	c.Read()
	fmt.Println("read action done!")
	c.Update()
	fmt.Println("update action done!")

}

//db related
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "test"
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
		fmt.Println("连接数据库错误：" + err.Error())
	} else {
		fmt.Println("Successfully connected!")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DBPing错误：" + err.Error())
	} else {
		fmt.Println("DBPingSuccess")
	}
	return &appContext{db}, ""
}

func lockDb() {
	return
}

func unlockDb() {
	return
}

func uploadS3() {
	return
}

// Create
func (c *appContext) Create() {
	// get insert id
	lastInsertId := 0
	now_str := time.Now().Format("2006-01-02 15:04:05")
	err := c.db.QueryRow("INSERT INTO userinfo(username,departname,Created) VALUES($1,$2,$3) RETURNING uid", "cruise", "软件部", now_str).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted id is ", lastInsertId)
}

// Read
func (c *appContext) Read() {
	rows, err := c.db.Query("SELECT * FROM userinfo ")
	//rows, err := c.db.Query("SELECT * FROM userinfo where uid = 3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := new(userinfo)
		err := rows.Scan(&p.uid, &p.username, &p.departname, &p.Created)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(p.uid, p.username, p.departname, p.Created)
	}
}

// UPDATE
func (c *appContext) Update() {
	stmt, err := c.db.Prepare("UPDATE userinfo SET username = $1 WHERE uid = $2")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec("jack", 1)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update affect rows is ", affectNum)
}

// DELETE
func (c *appContext) Delete() {
	stmt, err := c.db.Prepare("DELETE FROM userinfo WHERE uid = $1")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(1)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete affect rows is ", affectNum)
}
