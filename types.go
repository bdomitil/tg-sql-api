package main

type chat struct {
	ID         int    `json:"-" gorm:"column:id"`
	Chat_id    int64  `json:"chat_id" gorm:"column:chat_id"`
	Title      string `json:"title" gorm:"column:title"`
	Bot_id      int64  `json:"bot_id" gorm:"column:bot_id"`
	Type       uint8  `json:"type" gorm:"column:type"`
	Department string `json:"department" gorm:"column:department"`
}

type user struct {
	ID         int    `json:"-" gorm:"column:id"`
	User_id    int64  `json:"user_id" gorm:"column:user_id"`
	Bot_id     int64  `json:"bot_id" gorm:"column:bot_id"`
	Username   string `json:"username" gorm:"column:username"`
	Firstname  string `json:"firstname,omitempty" gorm:"column:firstname"`
	Rang       int8   `json:"rang" gorm:"column:rang"`
	Department string `json:"department" gorm:"column:department"`
}
