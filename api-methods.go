package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authChatHandler(c *gin.Context) {
	type response struct {
		Id   int  `json:"id"`
		Chat chat `json:"chat"`
	}
	var res response
	err := c.BindJSON(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
	} else { //Main logic

		sql, err := tg.OpenDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		}
		defer sql.Close()
		sqres, err := sql.Exec(`INSERT INTO tg_users (bot_id, chat_id, is_bot, chat_title)  VALUES (?, ?, ?, ?)`,
			res.Id, res.Chat.ID, res.Chat.IsBot, res.Chat.Title)
		if err != nil {
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

func listChatHandler(c *gin.Context) {
	type response struct {
		Id int `json:"id"`
	}
	var res response
	var chats []chat
	err := c.BindJSON(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	} else { //Main logic
		sql, err := tg.OpenDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
		defer sql.Close()
		rows, err := sql.Query("SELECT chat_id, is_bot, chat_title FROM tg_users WHERE bot_id = ?", res.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
		for rows.Next() {
			chat := chat{}
			err := rows.Scan(&chat.ID, &chat.IsBot, &chat.Title)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			chats = append(chats, chat)
		}
	}
	c.JSON(http.StatusOK, chats)
}
