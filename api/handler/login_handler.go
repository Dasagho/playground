package handler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/dasagho/playground/api/client"
	"github.com/dasagho/playground/api/external"
	"github.com/dasagho/playground/api/model"
	"github.com/gin-gonic/gin"
)

func PostLogin(c *gin.Context) {
	var login model.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}

	client := client.New()
	formData := login.FillLoginFormData()

	fmt.Printf("formData: %v\n", formData)
	loginResponse, err := client.PostForm("https://intranet.upv.es/pls/soalu/est_aute.intraalucomp", formData)
	if err != nil {
		panic("Error building login request" + err.Error())
	}
	defer loginResponse.Body.Close()

	if loginResponse.StatusCode != 200 {
		c.JSON(http.StatusBadRequest, gin.H{"Error on Post login request:": fmt.Sprintf("Status Code: %d", loginResponse.StatusCode)})
	}
	poliformatDocument, err := goquery.NewDocumentFromReader(loginResponse.Body)
	if err != nil {
		panic("Error parsing HTML response" + err.Error())
	}
	SubjectList := external.GetPoliformatMainPageData(*poliformatDocument, *client)
	c.JSON(http.StatusOK, SubjectList)
}
