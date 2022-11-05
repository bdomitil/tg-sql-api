package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//Initialization and testing connection to Mysql DB

var Db *gorm.DB

func Init() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	var err error
	db_name := os.Getenv("TG_DB")
	db_user := os.Getenv("TG_DB_USER")
	db_passwd := os.Getenv("TG_DB_PASSWD")
	db_host := os.Getenv("TG_DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_passwd, db_host, db_name)
	connector := mysql.Open(dsn)
	Db, err = gorm.Open(connector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Error : %s\n, %s\n", err, "Required: 'TG_DB, TG_DB_USER, TG_DB_PASSWD, TG_DB_HOST' variables")
	}
	log.Printf("Connected to DB %s successfully", db_name)
}

func main() {
	Init()
	router := gin.Default()
	gin.EnableJsonDecoderDisallowUnknownFields()
	chat := router.Group("/chat")
	{
		chat.GET("/list/", listAllChatsHandler)
		chat.GET("/list/:id", listChatHandler)
		chat.GET("/:bot_id/:chat_id/", getChatHandler)
		chat.POST("/add/", addChatHandler)
	}
	user := router.Group("/user")
	{
		user.GET("/:bot_id/:user_id/", getUserHandler)
		user.GET("/list/:bot_id", listUsersHandler)
		user.GET("/list/", listAllUsersHandler)
		user.POST("/add/", addUserHandler)

	}
	router.Run(":3334")
}
