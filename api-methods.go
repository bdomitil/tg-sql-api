package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addChatHandler(c *gin.Context) {
	var chat chat
	err := c.BindJSON(&chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}
	ch, err := getChatByBotID(chat.Bot_id, chat.Chat_id)
	if err != nil || ch.ID == 0 {
		err = addChat(chat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
	} else {
		err = updateChat(chat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, chat)
}

func listChatHandler(c *gin.Context) {
	ok, id := string_to_int64(c.Params.ByName("id"))
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"result": "Bad request id"})
		return
	}
	chats, err := getChatsByBotID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	if len(chats) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"result": "No such bot"})
		return
	}
	c.JSON(http.StatusOK, chats)
}

func getChatHandler(c *gin.Context) {
	ok, chat_id := string_to_int64(c.Params.ByName("chat_id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad id value"})
		return
	}
	ok, bot_id := string_to_int64(c.Params.ByName("bot_id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad id value"})
		return
	}
	chat, err := getChatByBotID(bot_id, chat_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chat)
}

func getUserHandler(c *gin.Context) {
	ok, user_id := string_to_int64(c.Params.ByName("user_id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad id value"})
		return
	}
	ok, bot_id := string_to_int64(c.Params.ByName("bot_id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad id value"})
		return
	}
	user, err := getUserByID(bot_id, user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func listUsersHandler(c *gin.Context) {
	ok, bot_id := string_to_int64(c.Params.ByName("bot_id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad BotId value"})
		return
	}

	users, err := getUsersByBotID(bot_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	if len(users) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"result": "No users for such botID"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func listAllUsersHandler(c *gin.Context) {

	users, err := getAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func listAllChatsHandler(c *gin.Context) {

	chats, err := getAllChats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chats)
}

func addUserHandler(c *gin.Context) {
	var user user
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	u, err := getUserByID(user.Bot_id, user.User_id)
	if u.ID == 0 || err != nil {
		err = addUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
	} else {
		log.Println("updating")
		err = updateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, user)
}

func deleteChatsHandler(c *gin.Context) {
	ok, chat_id := string_to_int64(c.Params.ByName("chat_id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad id value"})
		return
	}
	ok, bot_id := string_to_int64(c.Params.ByName("bot_id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"result": "Bad id value"})
		return
	}
	status, err := deleteChat(chat{Bot_id: bot_id, Chat_id: chat_id})
	switch status {
	case 500:
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	case 200:
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	default:
		c.JSON(http.StatusNotFound, gin.H{"result": "fail"})
	}

}
