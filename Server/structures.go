package main

// import "golang.org/x/text/date"

// import ("go.mongodb.org/mongo-driver/bson")

type Event struct {
	ID string `bson:"id"`
	Name string `bson:"name"`
	Desc string `bson:"desc"`
	RDate string `bson:"date"`
}

type Auth struct {
	Username string `bson:"uname"`
	Pass string `bson:"pass"`
	Salt string `bson:"salt"`
}