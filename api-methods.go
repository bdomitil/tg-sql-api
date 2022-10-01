package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addChatHandler(c *gin.Context) {
	var chat chat
	err := c.BindJSON(&chat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
	} else { //Main logic

		sql, err := tg.OpenDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		}
		defer sql.Close()
		sqlcheck, err := sql.Query("SELECT * FROM chat WHERE bot = ? and id = ?", chat.BotID, chat.ID)
		if err != nil || sqlcheck.Next() {
			c.JSON(http.StatusOK, gin.H{"result": "Bot is already registered"})
			return
		} else {
			sqres, err := sql.Exec(`INSERT INTO chat (id, title, bot, type)  VALUES (?, ?, ?, ?)`,
				chat.ID, chat.Title, chat.BotID, chat.Type)
			if err != nil {
				fmt.Println(err.Error())

				c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
				return
			}
			if n, err := sqres.RowsAffected(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
				return
			} else if n == 0 {
				c.JSON(http.StatusOK, gin.H{"result": "Already exists"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"result": "Successfuly added"})
		}
	}
}

func listChatHandler(c *gin.Context) {

	var chats []chat
	ok, id := string_to_int64(c.Params.ByName("id"))

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"result": "Bad request id"})
		return
	}
	//Main logic
	sql, err := tg.OpenDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	defer sql.Close()
	rows, err := sql.Query("SELECT id, title, type FROM chat WHERE bot = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	} else if !rows.Next() {
		c.JSON(http.StatusNotFound, gin.H{"result": "No such bot"})
		return
	}
	for rows.Next() {
		chat := chat{}
		err := rows.Scan(&chat.ID, &chat.Title, &chat.Type)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		chat.BotID = id
		chats = append(chats, chat)
	}
	c.JSON(http.StatusOK, chats)
}

func getChatHandler(c *gin.Context) {

}

func getUserHandler(c *gin.Context) {
	var reqUser user
	ok, id := string_to_int64(c.Params.ByName("id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad id value"})
		return
	}
	sql, err := tg.OpenDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	defer sql.Close()
	rows, err := sql.Query("SELECT username, firstname, rang, department FROM user WHERE id = ?", id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	if !rows.Next() {
		c.JSON(http.StatusNotFound, gin.H{"result": "No such user"})
		return
	}
	err = rows.Scan(&reqUser.Username, &reqUser.Firstname, &reqUser.Rang, &reqUser.Department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	reqUser.ID = id
	c.JSON(http.StatusOK, reqUser)
}

func listUserHandler(c *gin.Context) {
	var reqUsers []user
	ok, id := string_to_int64(c.Params.ByName("botID"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad BotId value"})
		return
	}
	sql, err := tg.OpenDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	defer sql.Close()

	rows, err := sql.Query("SELECT id, username, firstname, rang, department FROM user WHERE botId = ? ", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	if !rows.Next() {
		c.JSON(http.StatusNotFound, gin.H{"result": "No users for such botID"})
		return
	}
	for rows.Next() {
		newUser := user{}
		err = rows.Scan(&newUser.ID, &newUser.Username, &newUser.Firstname, &newUser.Rang, &newUser.Department)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
		newUser.BotID = id
		reqUsers = append(reqUsers, newUser)
	}
	c.JSON(http.StatusOK, reqUsers)
}

func addUserHandler(c *gin.Context) {
	var user user
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	sql, err := tg.OpenDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	defer sql.Close()
	sqlres, err := sql.Exec(`INSERT INTO user (id, botId, username, firstname, rang, department) values (?, ?, ?, ?, ?, ?) `, user.ID, user.BotID, user.Username, user.Firstname, user.Rang, user.Department) //TODO check if all necessary fields present
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	af, _ := sqlres.RowsAffected()
	if af < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}
