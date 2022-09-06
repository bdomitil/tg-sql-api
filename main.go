package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//Initialization and testing connection to Mysql DB

var tg tg_cache

func xinit(db_name, login, password, hostname string) error {
	tg.db.login = login
	tg.db.name = db_name
	tg.db.password = password
	tg.db.hostname = hostname
	sql, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", tg.db.login, tg.db.password, tg.db.hostname, tg.db.name))
	if err != nil {
		panic(err)
	}
	log.Printf("Connected to DB %s successfully", db_name)
	defer sql.Close()
	return err
}

func (tg *tg_cache) OpenDB() (*sql.DB, error) {
	sql, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", tg.db.login, tg.db.password, tg.db.hostname, tg.db.name))
	if err != nil {
		log.Println(err)
	}
	return sql, err
}

func main() {
	db_name := os.Getenv("TG_DB")
	db_user := os.Getenv("TG_DB_USER")
	db_passwd := os.Getenv("TG_DB_PASSWD")
	db_host := os.Getenv("TG_DB_HOST")
	if db_name == "" || db_user == "" || db_passwd == "" || db_host == "" {
		panic("Required: 'TG_DB, TG_DB_USER, TG_DB_PASSWD, TG_DB_HOST' variables")
	}
	xinit(db_name, db_user, db_passwd, db_host)
	router := gin.Default()
	chat := router.Group("/chat")
	{
		chat.POST("/add/", addChatHandler)
		chat.POST("/list/", listChatHandler)
	}
	router.Run(":3334")
}
