package main

import (
	"net/http"
)

func (app *App) Authenticate(w http.ResponseWriter, r http.Request) (User, error) {
	var user User
	err := r.ParseForm()
	if err != nil {
		return user, err
	}
	token := r.PostForm.Get("token")
	app.db.Where(&User{
		Token: token,
	}).Take(&user)
	return user, nil
}
