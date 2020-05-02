package service

import (
	"github.com/personal-website-go/dao"
	"github.com/personal-website-go/models"
)

type ContentService struct {
	dao.BaseDao
}

func (i *ContentService) Init(contentDao dao.BaseDao) {
	i.BaseDao = contentDao
}

func (i *ContentService) Create(content models.Content) (err error) {
	return i.BaseDao.Create(content)
}

func (i *ContentService) Update(id string, content models.Content) (err error) {
	selector, err := i.BaseDao.FindById(id)
	if err != nil {
		return
	}
	return i.BaseDao.Update(selector, content)
}

func (i *ContentService) FindById(id string) (content models.Content, err error) {
	return i.BaseDao.FindById(id)
}
