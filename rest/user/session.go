package user

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	sessionCookieName = "blackbookSession"
)

func newSessionCookie(id uint, session string) *http.Cookie {
	return &http.Cookie{
		Name:     sessionCookieName,
		Value:    fmt.Sprintf("%d:%s", id, session),
		MaxAge:   2592000,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
}

func parseSessionCookie(cookie *http.Cookie) (uint, string, error) {
	var (
		id64    uint64
		session string
		err     error
	)
	parts := strings.Split(cookie.Value, ":")

	if len(parts) != 2 {
		return uint(id64), session, fmt.Errorf("could not separate cookie")
	}

	session = parts[1]

	id64, err = strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return uint(id64), session, fmt.Errorf("could not parse id in cookie")
	}

	return uint(id64), session, err
}

func findSessionCookie(r *http.Request) *http.Cookie {
	for _, cookie := range r.Cookies() {
		if cookie.Name == sessionCookieName {
			return cookie
		}
	}

	return nil
}
