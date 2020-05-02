package service

import (
	"github.com/personal-website-go/dao"
	"github.com/personal-website-go/models"
)

type BaseService interface {
	Init(contentDao dao.BaseDao)
	Create(content models.Content) (err error)
	Update(id string, content models.Content) (err error)
	FindById(id string) (content models.Content, err error)
}
