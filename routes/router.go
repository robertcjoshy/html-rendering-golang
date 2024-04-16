package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/html-rendering-golang/handlers"
)

func Startroutes(r *gin.Engine) {
	r.GET("/", handlers.Indexpage)
	r.GET("/about", handlers.Aboutpage)
	r.Run()
}
