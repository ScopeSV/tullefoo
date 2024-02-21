package main

import (
	foocontroller "github.com/ScopeSV/tulle-repo/internal/controllers/foo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fooController := foocontroller.NewFooController()

	r.GET("/foo", fooController.GetFoo)
}
