package model

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/dasagho/playground/api/client"
)

type User struct {
	Id         string        `json:"id"`         // Cookie provided by us
	Tdp        http.Cookie   `json:"tdp"`        // Both cookies provided by Poliformat
	JsessionID http.Cookie   `json:"jsessionId"` // and by Upv
	Subjects   []Subject     `json:"subjects"`   // Slice of subjects
	client     client.Client // Cliente http con las cookies del usuario
}

func NewUser(id string, tdp, jsession http.Cookie, subjects []Subject, client client.Client) User {
	return User{
		Id:         id,
		Tdp:        tdp,
		JsessionID: jsession,
		Subjects:   subjects,
		client:     client,
	}
}

func (u User) getCSRFToken(subject *Subject) {
	resourcesPage, err := u.client.Get(subject.Recursos)
	if err != nil {
		panic("Error request resources page")
	}
	defer resourcesPage.Body.Close()

	resourcesDoc, err := goquery.NewDocumentFromReader(resourcesPage.Body)
	if err != nil {
		panic("Error parsing resources HTML")
	}

	name, existName := resourcesDoc.Find("[name='sakai_csrf_token']").First().Attr("value")
	if !existName {
		return
	}
	subject.CSRF = name
}
