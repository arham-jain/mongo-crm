package controller

import (
	"encoding/json"
	"net/http"

	"github.com/personal-website-go/dep"
	"github.com/personal-website-go/models"
)

type ContentController struct {
	BaseController
}

// @router / [post]
// @Accept json
func (i *ContentController) Create() {
	var c models.Content
	err := json.Unmarshal(i.Ctx.Input.RequestBody, &c)
	if err != nil {
		i.serve(http.StatusBadRequest, nil)
	}
	err = dep.ContentService.Create(c)
	if err != nil {
		i.serve(http.StatusInternalServerError, nil)
	}
	i.serve(http.StatusCreated, c)
}

// @router /:id [put]
// @Accept json
func (i *ContentController) Update() {
	var c models.Content
	err := json.Unmarshal(i.Ctx.Input.RequestBody, &c)
	if err != nil {
		i.serve(http.StatusBadRequest, nil)
	}
	id := i.Ctx.Input.Param(":id")
	err = dep.ContentService.Update(id, c)
	if err != nil {
		i.serve(http.StatusInternalServerError, nil)
	}
	i.serve(http.StatusCreated, c)
}

// @router /:id [get]
func (i *ContentController) FindById() {
	id := i.Ctx.Input.Param(":id")
	c, err := dep.ContentService.FindById(id)
	if err != nil {
		i.serve(http.StatusNotFound, nil)
	}
	i.serve(http.StatusOK, c)
}
