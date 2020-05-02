package models

type Body struct {
	Title 		string	`bson:"title"`
	Content 	string	`bson:"content"`
	Keywords    string	`bson:"keyword"`
	Language	string	`bson:"language"`
}