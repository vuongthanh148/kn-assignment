// Package cookie is used to get and set HTTP cookies.
package cookie

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/kwanpham2195/go-gcp-boilerplate/internal/derrors"
)

// Extract returns the value of the cookie at name and deletes the cookie.
func Extract(w http.ResponseWriter, r *http.Request, name string) (_ string, err error) {
	defer derrors.Wrap(&err, "Extract")
	c, err := r.Cookie(name)
	if err != nil && err != http.ErrNoCookie {
		return "", fmt.Errorf("r.Cookie(%q): %v", name, err)
	}
	if c == nil {
		return "", nil
	}
	val, err := Base64Value(c)
	if err != nil {
		return "", err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Path:    r.URL.Path,
		Expires: time.Unix(0, 0),
	})
	return val, nil
}

// Base64Value decodes  the value of c using the Base64 URL encoding and returns it as a string.
func Base64Value(c *http.Cookie) (string, error) {
	val, err := base64.URLEncoding.DecodeString(c.Value)
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// Set sets a cookie at the urlPath with name and val.
func Set(w http.ResponseWriter, name, val, urlPath string) {
	value := base64.URLEncoding.EncodeToString([]byte(val))
	http.SetCookie(w, &http.Cookie{Name: name, Value: value, Path: urlPath})
}
