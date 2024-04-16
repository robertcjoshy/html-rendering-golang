package main

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robert/html-rendering-golang/routes"
)

func main() {

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")

	routes.Startroutes(r)

}
