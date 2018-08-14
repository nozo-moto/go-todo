package controller

import (
	"net/http"
)

var userdata = []User{
	{
		"hoge",
	},
	{
		"fuga",
	},
	{
		"moge",
	},
}

type User struct {
	Name string `json:name`
}

func (u *User) Get(w http.ResponseWriter, r *http.Request) error {
	return JSON(w, http.StatusOK, userdata)
}
