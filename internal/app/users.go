package app

import (
	"net/http"

	"github.com/kevindurb/togo/internal/database"
	"github.com/kevindurb/togo/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserBody struct {
	Username string `schema:"username,required"`
	Password string `schema:"password,required"`
}

type NewUserPageData struct {
	Err string
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body CreateUserBody

	if err := utils.DecodePostForm(&body, r, a.decoder); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	a.queries.CreateUser(r.Context(), database.CreateUserParams{
		Username:     body.Username,
		PasswordHash: hash,
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (a *App) NewUser(w http.ResponseWriter, r *http.Request) {
	s := getAppSession(r)
	err := getFlashError(s)
	s.Save(r, w)

	a.newUserTmpl.ExecuteTemplate(w, "base", NewUserPageData{
		Err: err,
	})
}
