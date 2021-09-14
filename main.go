package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type contact struct {
	ID     string
	Name   string
	Number string
	Email  string
}

var contacts = []contact{
	{ID: "1", Name: "foo bar", Number: "(82) 9999-1234", Email: "foo@bar.com"},
	{ID: "2", Name: "ping pong", Number: "(82) 9999-4567", Email: "ping@pong.com"},
	{ID: "3", Name: "zig zag", Number: "(82) 9999-6789", Email: "zig@zag.com"},
}

func getContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

func main() {
	r := gin.Default()

	r.GET("/contacts", getContacts)

	r.Run(":3000")
}
