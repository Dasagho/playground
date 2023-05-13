package services

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/dasagho/playground/api/client"
	"github.com/dasagho/playground/api/model"
	"github.com/dasagho/playground/external"
)

const loginEndpoint string = "https://intranet.upv.es/pls/soalu/est_aute.intraalucomp"
const poliformatURL string = "https://poliformat.upv.es"
const upvURL string = "https://upv.es"

func GetUserClient(login model.Login) *client.Client {
	client := client.New()
	return client
}

func LoginUser(login model.Login, client *client.Client) (*http.Response, error) {
	formData := login.FillLoginFormData()
	loginResponse, err := client.PostForm(loginEndpoint, formData)
	if err != nil {
		return nil, errors.New("Error building login request " + err.Error())
	}

	if loginResponse.StatusCode != 200 {
		return nil, errors.New("Error on Post login request")
	}
	return loginResponse, nil
}

func GetSubjects(login model.Login, client *client.Client, loginResponse *http.Response) ([]model.Subject, error) {
	poliformatDocument, err := goquery.NewDocumentFromReader(loginResponse.Body)
	if err != nil {
		return nil, errors.New("Error parsing HTML response " + err.Error())
	}
	defer loginResponse.Body.Close()
	return external.GetPoliformatMainPageData(*poliformatDocument, *client), nil
}

func GetUser(userList []model.User, client *client.Client, SubjectList []model.Subject) (model.User, error) {
	poliformat, err := url.Parse(poliformatURL)
	if err != nil {
		return model.User{}, errors.New("Error searching poliformat cookie " + err.Error())
	}

	upv, err := url.Parse(upvURL)
	if err != nil {
		return model.User{}, errors.New("Error searching upv cookie " + err.Error())
	}

	return model.NewUser(
		len(userList),
		*client.Jar.Cookies(poliformat)[1],
		*client.Jar.Cookies(upv)[0],
		SubjectList,
		*client,
	), nil
}

func FindUser(id int, userList []model.User) (model.User, error) {
	var userSearched model.User
	for _, user := range userList {
		if user.Id == id {
			userSearched = user
			break
		}
	}

	if userSearched.JsessionID.Name == "" {
		return model.User{}, errors.New(fmt.Sprintf("Can't find User with Id: %d", id))
	}

	return userSearched, nil
}
