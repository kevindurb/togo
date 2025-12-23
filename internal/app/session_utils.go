package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func (a *App) ProtectRoute(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := getAppSession(r)

		sessID := getSessID(s)
		session, err := a.queries.GetSession(r.Context(), sessID)

		s.Save(r, w)

		if err != nil || time.Now().After(session.ExpiresAt) {
			log.Printf("Session (%s) error: %s", sessID, err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		handler(
			w,
			r.WithContext(
				context.WithValue(r.Context(), "userID", session.UserID),
			),
		)
	}
}

func getSessID(s *sessions.Session) string {
	if sessID, ok := s.Values["sessID"].(string); ok {
		return sessID
	}

	return ""
}

func saveSessID(w http.ResponseWriter, r *http.Request, sessID uuid.UUID) {
	s := getAppSession(r)
	s.Values["sessID"] = sessID.String()
	s.Save(r, w)
}

func getFlashError(s *sessions.Session) string {
	flashes := s.Flashes("error")
	if len(flashes) == 0 {
		return ""
	}

	if msg, ok := flashes[0].(string); ok {
		return msg
	}

	return ""
}

func getAppSession(r *http.Request) *sessions.Session {
	s, err := sessionStore.Get(r, "togo-session")
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func userIDFromContext(c context.Context) int64 {
	return c.Value("userID").(int64)
}
