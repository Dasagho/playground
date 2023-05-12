package model

import (
	"net/url"
	"strconv"
)

type Login struct {
	User int    `json:"user"`
	Pass string `json:"pass"`
}

func (l Login) FillLoginFormData() url.Values {
	dni := l.User
	pass := l.Pass

	formData := url.Values{}
	formData.Add("id", "c")
	formData.Add("estilo", "500")
	formData.Add("vista", "MSE")
	formData.Add("cua", "sakai")
	formData.Add("dni", strconv.Itoa(dni))
	formData.Add("clau", pass)

	return formData
}
