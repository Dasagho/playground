package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/dasagho/playground/api/client"
	"github.com/dasagho/playground/api/external"
	"github.com/dasagho/playground/api/model"
	"github.com/gin-gonic/gin"
)

var userList []model.User

const loginEndpoint string = "https://intranet.upv.es/pls/soalu/est_aute.intraalucomp"
const poliformatURL string = "https://poliformat.upv.es"
const upvURL string = "https://upv.es"

func PostLogin(c *gin.Context) {
	var login model.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}

	client := client.New()
	formData := login.FillLoginFormData()

	loginResponse, err := client.PostForm(loginEndpoint, formData)
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

	poliformat, _ := url.Parse(poliformatURL)
	upv, _ := url.Parse(upvURL)

	user := model.NewUser(
		strconv.Itoa(len(userList)),
		*client.Jar.Cookies(poliformat)[0],
		*client.Jar.Cookies(upv)[0],
		SubjectList,
		*client,
	)

	userList = append(userList, user)
	c.JSON(http.StatusOK, SubjectList)
}

func GetUsersLoged(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}
