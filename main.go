package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type contact struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	Email  string `json:"email,omitempty"`
}

var contacts = []contact{
	{ID: "1", Name: "foo bar", Number: "(82) 9999-1234", Email: "foo@bar.com"},
	{ID: "2", Name: "ping pong", Number: "(82) 9999-4567", Email: "ping@pong.com"},
	{ID: "3", Name: "zig zag", Number: "(82) 9999-6789", Email: "zig@zag.com"},
}

func getContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

func postContacts(c *gin.Context) {
	var newContact contact

	if err := c.BindJSON(&newContact); err != nil {
		return
	}

	contacts = append(contacts, newContact)
	c.IndentedJSON(http.StatusCreated, newContact)
}

func getContact(c *gin.Context) {
	id := c.Param("id")

	for _, contact := range contacts {
		if contact.ID == id {
			c.IndentedJSON(http.StatusOK, contact)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contact not found"})
}

func main() {
	r := gin.Default()

	r.GET("/contacts", getContacts)
	r.GET("/contacts/:id", getContact)
	r.POST("/contacts", postContacts)

	r.Run(":3000")
}
