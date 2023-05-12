package model

import (
	"net/url"
	"strings"
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
