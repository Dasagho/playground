package handler

import (
	"net/http"

	"github.com/dasagho/playground/api/model"
	"github.com/dasagho/playground/api/services"
	"github.com/gin-gonic/gin"
)

var userList []model.User

func PostLogin(c *gin.Context) {
	var login model.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
	}

	client := services.GetUserClient(login)
	res, err := services.LoginUser(login, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	subjectList, err := services.GetSubjects(login, client, res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	user, err := services.GetUser(userList, client, subjectList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	userList = append(userList, user)
	c.JSON(http.StatusOK, subjectList)
}
