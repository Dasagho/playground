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

func SendLoginRequest(login model.Login, client *client.Client) (*http.Response, error) {
	formData := login.FillLoginFormData()

	loginResponse, err := client.PostForm(loginEndpoint, formData)

	if err != nil {
		return nil, errors.New("Error building login request " + err.Error())
	}

	return loginResponse, nil
}

func CheckLogin(login model.Login, loginResponse *http.Response) (goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(loginResponse.Body)

	defer loginResponse.Body.Close()

	if err != nil {
		return goquery.Document{}, fmt.Errorf("error parsing poliformat login response %s", err.Error())
	}

	dniString := doc.Find(".Mrphs-userNav__submenuitem--userid").First().Text()
	if len(dniString) == 0 {
		return goquery.Document{}, fmt.Errorf("error login failed")
	}

	// dni, err := strconv.Atoi(dniString)
	// if err != nil {
	// 	return goquery.Document{}, fmt.Errorf("error casting string to int %s", err.Error())
	// }

	// if dni != login.User {
	// 	return goquery.Document{}, fmt.Errorf("error dni mismatch %d != %d", dni, login.User)
	// }

	return *doc, nil
}

func GetSubjects(login model.Login, client *client.Client, loginResponse goquery.Document) []model.Subject {
	return external.GetPoliformatMainPageData(loginResponse, *client)
}

func CreateUser(userList []model.User, client *client.Client, SubjectList []model.Subject) (model.User, error) {
	poliformat, err := url.Parse(poliformatURL)
	if err != nil {
		return model.User{}, errors.New("Error searching poliformat cookie " + err.Error())
	}

	upv, err := url.Parse(upvURL)
	if err != nil {
		return model.User{}, errors.New("Error searching upv cookie " + err.Error())
	}

	poliCookies := client.Jar.Cookies(poliformat)
	if len(poliCookies) < 1 {
		return model.User{}, fmt.Errorf("failed find cookies for domain: %s", poliformat)
	}

	upvCookies := client.Jar.Cookies(upv)
	if len(upvCookies) < 1 {
		return model.User{}, fmt.Errorf("failed find cookies for domain: %s", upv)
	}

	jsessionId := poliCookies[1]
	tpd := upvCookies[0]

	return model.NewUser(
		len(userList),
		*jsessionId,
		*tpd,
		SubjectList,
		*client,
	), nil
}
