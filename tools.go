package main

import (
	"errors"
	"math/big"
)

func string_to_int64(str string) (bool, int64) {

	big, ok := new(big.Int).SetString(str, 10)
	if ok {
		return ok, big.Int64()
	} else {
		return ok, 0
	}
}

func getAllChats() (Chats []chat, err error) {
	result := Db.Find(&Chats)
	err = result.Error
	return
}

func getChatsByBotID(botID int64) (Chats []chat, err error) {
	result := Db.Where("bot_id = ?", botID).Find(&Chats)
	err = result.Error
	return
}

func getChatByBotID(bot_id int64, chat_id int64) (Chat chat, err error) {
	result := Db.Where("bot_id = ? and chat_id = ?", bot_id, chat_id).Take(&Chat)
	if result.Error != nil {
		err = result.Error
	} else if result.RowsAffected < 0 {
		err = errors.New("no such chat")
	}
	return
}

func addChat(Chat chat) error {
	result := Db.Create(&Chat)
	return result.Error
}

func updateChat(Chat chat) error {
	result := Db.Where("chat_id = ? and bot_id = ?", Chat.Chat_id, Chat.Bot_id).Updates(Chat)
	return result.Error
}

func getUserByID(bot_id, user_id int64) (User user, err error) {
	result := Db.Where("user_id = ? and bot_id = ?", user_id, bot_id).Take(&User)
	if result.Error != nil {
		err = result.Error
	} else if result.RowsAffected < 0 {
		err = errors.New("no such user")
	}
	return
}

func getAllUsers() (Users []user, err error) {
	result := Db.Find(&Users)
	err = result.Error
	return
}

func getUsersByBotID(bot_id int64) (Users []user, err error) {
	result := Db.Where("bot_id = ?", bot_id).Take(&Users)
	err = result.Error
	return
}

func addUser(User user) error {
	result := Db.Create(&User)
	return result.Error
}

func updateUser(User user) error {
	result := Db.Where("bot_id = ? and user_id = ?", User.Bot_id, User.User_id).Updates(User)
	return result.Error
}
