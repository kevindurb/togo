package app

import (
	"io/fs"
	"net/http"

	"github.com/kevindurb/togo/web"
)

func (a *App) StaticFileServer() http.Handler {
	staticFs, _ := fs.Sub(web.Files, "static")
	return http.FileServer(http.FS(staticFs))
}
