package handler

import (
	"net/http"

	"github.com/dasagho/playground/api/model"
	"github.com/dasagho/playground/api/services"
	"github.com/gin-gonic/gin"
)

var userList model.UserList

func PostLogin(c *gin.Context) {
	var login model.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}

	client := services.GetUserClient(login)
	res, err := services.LoginUser(login, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	poliPage, err := services.CheckLogin(login, res)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return
	}

	subjectList := services.GetSubjects(login, client, poliPage)

	user, err := services.CreateUser(userList.Get(), client, subjectList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	userList.Append(user)

	c.JSON(http.StatusOK, subjectList)
}
