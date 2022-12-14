package main

type tg_db struct {
	name     string
	login    string
	password string
	hostname string
}

type tg_cache struct {
	db       tg_db
	login    string
	password string
}

type chat struct {
	ID         int64  `json:"chat_id"`
	Title      string `json:"title"`
	BotID      int64  `json:"botID"`
	Type       uint8  `json:"type"`
	Department string `json:"department"`
}

type user struct {
	ID         int64  `json:"user_id"`
	BotID      int64  `json:"botId"`
	Username   string `json:"username"`
	Firstname  string `json:"firstname,omitempty"`
	Rang       int8   `json:"rang"`
	Department string `json:"department"`
}

type taskStore interface {
	getStore() *taskStore
}

func (*user) getStore() *user {
	sUser := &user{}
	return sUser
}

func (*chat) getStore() *chat {
	sChat := &chat{}
	return sChat
}
