package handler

import (
	"net/http"
	"strconv"

	"github.com/dasagho/playground/api/services"
	"github.com/gin-gonic/gin"
)

func GetResources(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	user, err := services.FindUser(id, userList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, user.Subjects)
}
