package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//Initialization and testing connection to Mysql DB

var tg tg_cache

func xinit(db_name, login, password, address, port string) error {
	tg.db.login = login
	tg.db.name = db_name
	tg.db.password = password
	tg.db.hostname = address + ":" + port
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
	xinit("telegram_bitrix", "bitrix_bot", "Kt)J@58m2u", "127.0.0.1", "3306")
	router := gin.Default()
	chat := router.Group("/chat")
	{
		chat.POST("/auth/", authChatHandler)
		chat.POST("/list/", listChatHandler)
	}
	router.Run(":3334")
}
