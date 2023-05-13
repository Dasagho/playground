package model

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dasagho/playground/api/client"
)

type Subject struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Recursos string `json:"recursos"`
	CSRF     string `json:"csrf"`
}

func (s Subject) fillResourcesFormData() url.Values {
	collectionId := "/group/" + strings.Split(s.Link, "/")[4] + "/"
	formData := url.Values{}
	formData.Add("sakai_action", "doExpandall")
	formData.Add("collectionid", collectionId)
	formData.Add("sakai_csrf_token", s.CSRF)
	return formData
}

func (s Subject) getCSRFToken(client client.Client, subject *Subject) {
	resourcesPage, err := client.Get(subject.Recursos)
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
