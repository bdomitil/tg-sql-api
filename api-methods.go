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
		sqlcheck, err := sql.Query("SELECT id FROM chat WHERE bot = ?", chat.BotID)
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
		rows, err := sql.Query("SELECT id, title FROM tg_users WHERE bot = ?", res.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
		for rows.Next() {
			chat := chat{}
			err := rows.Scan(&chat.ID, &chat.Title)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			chats = append(chats, chat)
		}
	}
	c.JSON(http.StatusOK, chats)
}
