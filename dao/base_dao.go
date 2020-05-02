package dao

import (
	"github.com/go-bongo/bongo"
	"github.com/personal-website-go/models"
)

type BaseDao interface {
	Init(dbConn *bongo.Connection)
	Create(content models.Content) (err error)
	Update(selector models.Content, content models.Content) (err error)
	FindById(id string) (content models.Content, err error)
}