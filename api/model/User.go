package model

import (
	"net/http"

	"github.com/dasagho/playground/api/client"
)

type User struct {
	Id         int           `json:"id"`         // Cookie provided by us
	Tdp        http.Cookie   `json:"tdp"`        // Both cookies provided by Poliformat
	JsessionID http.Cookie   `json:"jsessionId"` // and by Upv
	Subjects   []Subject     `json:"subjects"`   // Slice of subjects
	client     client.Client // Cliente http con las cookies del usuario
}

func NewUser(id int, tdp, jsession http.Cookie, subjects []Subject, client client.Client) User {
	return User{
		Id:         id,
		Tdp:        tdp,
		JsessionID: jsession,
		Subjects:   subjects,
		client:     client,
	}
}
