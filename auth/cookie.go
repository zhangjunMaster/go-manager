package auth

import (
	"net/http"
)

/**
// http.Cookie
type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

*/

func Read(w http.ResponseWriter, req *http.Request) (string, error) {
	c, err := req.Cookie("manager")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return "", err
	}
	return c.Value, nil
}

func Set(w http.ResponseWriter, req *http.Request, value string) {
	cookie := &http.Cookie{
		Name:     "manager",
		Value:    value,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   30 * 60,
	}
	http.SetCookie(w, cookie)
}
