package ports

import "net/http"

type CookieSetter interface {
	SetCookie(*http.Cookie)
}

type CookieExtractor interface {
	Cookie(string) (*http.Cookie, error)
}

type HeaderSetter interface {
	Set(string, string)
}

type HeaderExtractor interface {
	Get(string) string
}

type Header interface {
	HeaderSetter
	HeaderExtractor
}
