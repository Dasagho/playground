package handler

import (
	"net/http"
	"strconv"

	"github.com/dasagho/playground/api/model"
	"github.com/gin-gonic/gin"
)

func GetUsersLoged(c *gin.Context) {
	c.JSON(http.StatusOK, userList.Get())
}

func GetUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	userSearched, err := userList.Find(func(u model.User) bool {
		return u.Id == id
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userSearched)
}
