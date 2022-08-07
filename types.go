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
	ID    int64  `json:"id"`
	IsBot bool   `json:"isBot"`
	Title string `json:"title"`
}

type user struct {
	ID        int64  `json:"id"`
	FirstName string `json:"name"`
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
