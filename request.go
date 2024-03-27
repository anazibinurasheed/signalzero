package main

type User struct {
	UserID   string `json:"user_id" bson:"user_id,omitempty" `
	UserName string `json:"username" bson:"username,omitempty" `
	Password string `json:"password" bson:"password,omitempty" `
}


