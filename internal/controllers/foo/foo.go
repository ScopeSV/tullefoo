package foocontroller

import (
	"net/http"

	foomodel "github.com/ScopeSV/tulle-repo/internal/models/foo"
	"github.com/gin-gonic/gin"
)

type FooController struct {
	model *foomodel.FooModel
}

func (fc *FooController) GetFoo(c *gin.Context) {
	collectedFoo := fc.model.GetFoo()

	c.JSON(http.StatusOK, gin.H{
		"result": collectedFoo,
	})
}

func NewFooController() *FooController {
	return &FooController{
		model: foomodel.NewFooModel(),
	}
}
