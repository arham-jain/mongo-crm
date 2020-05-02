package dao

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
	"github.com/personal-website-go/models"
)

type ContentDao struct {
	*bongo.Connection
}

func (i ContentDao) Init(dbConn *bongo.Connection) {
	i.Connection = dbConn
}

func (i ContentDao) Create(content models.Content) (err error) {
	err = i.Collection(content.CollectionName()).Save(content)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		fmt.Println("Validation errors :", vErr.Errors)
	} else {
		fmt.Println("Error in creating document :", err.Error())
	}
	return
}

func (i ContentDao) Update(selector models.Content, content models.Content) (err error) {
	err = i.Collection(content.CollectionName()).Collection().Update(selector, content)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		fmt.Println("Validation errors :", vErr.Errors)
	} else {
		fmt.Println("Error in updating document :", err.Error())
	}
	return
}

func (i ContentDao) FindById(id string) (content models.Content, err error) {
	err = i.Collection(content.CollectionName()).FindById(bson.ObjectIdHex(id), content)
	if err != nil {
		fmt.Println("Error in finding document :", err.Error())
	}
	return
}
