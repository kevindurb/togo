package utils

import (
	"net/http"

	"github.com/gorilla/schema"
)

func DecodePostForm(dst any, r *http.Request, d *schema.Decoder) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return d.Decode(&dst, r.PostForm)
}
