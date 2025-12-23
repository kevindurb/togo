package app

import (
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/kevindurb/togo/internal/database"
	"golang.org/x/crypto/bcrypt"
)

var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type LoginPageData struct {
	Err string
}

type LoginBody struct {
	Username string `schema:"username,required"`
	Password string `schema:"password,required"`
}

func showLoginError(w http.ResponseWriter, r *http.Request, msg string) {
	s := getAppSession(r)
	s.AddFlash(msg, "error")
	s.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		showLoginError(w, r, "Login Error")
		return
	}

	var body LoginBody
	err := a.decoder.Decode(&body, r.PostForm)
	if err != nil {
		showLoginError(w, r, "Login Error")
		return
	}

	user, err := a.queries.GetUser(r.Context(), body.Username)
	if err != nil {
		showLoginError(w, r, "Incorrect Username or Password")
		return
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(body.Password))
	if err != nil {
		showLoginError(w, r, "Incorrect Username or Password")
		return
	}

	sessID := uuid.New()
	expiresAt := time.Now().AddDate(0, 0, 30)

	err = a.queries.CreateSession(r.Context(), database.CreateSessionParams{
		ID:        sessID,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		showLoginError(w, r, "Login Error")
		return
	}
	saveSessID(w, r, sessID)

	http.Redirect(w, r, "/todos", http.StatusFound)
}

func (a *App) ShowLogin(w http.ResponseWriter, r *http.Request) {
	s := getAppSession(r)
	err := getFlashError(s)
	s.Save(r, w)

	a.showLoginTmpl.ExecuteTemplate(w, "base", LoginPageData{
		Err: err,
	})
}
