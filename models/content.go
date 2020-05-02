package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

// Content database structure
type Content struct {
	bongo.DocumentBase `bson:",inline"`
	Body               []Body  `json:"body"`
	Images             []Image `bson:"image"`
	Status             bool    `bson:"status"`
}

func (i *Content) CollectionName() string {
	return "content"
}

func (i Content) GetId() bson.ObjectId {
	return i.Id
}
func (i Content) SetId(ID bson.ObjectId) {
	i.Id = ID
}