package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Content string
	Ronot   string
}

func Indexpage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "this is a index page",
	})
}

// func Aboutpage(c *gin.Context) {
// 	c.HTML(http.StatusOK, "about.html", gin.H{
// 		"content": "this is a about page",
// 		"ronot":   "this is second data",
// 	})
// }

func Aboutpage(c *gin.Context) {

	value := data{
		Content: "robert",
		Ronot:   "messi",
	}
	slices := make([]data, 0)

	slices = append(slices, value)

	c.HTML(http.StatusOK, "about.html", gin.H{
		"slices": slices,
	})
}
