package handlers

import (
	"net/url"
	"strings"
)

const NextParam = "next"

func SafeNext(target string) string {
	if target == "" {
		return "/"
	}

	u, err := url.Parse(target)
	if err != nil {
		return "/"
	}

	if u.Scheme != "" || u.Host != "" || u.User != nil {
		return "/"
	}

	if !strings.HasPrefix(u.Path, "/") || strings.HasPrefix(u.Path, "//") {
		return "/"
	}

	if strings.ContainsAny(u.Path, "\\\r\n") {
		return "/"
	}

	if strings.IndexFunc(u.Path, func(r rune) bool { return r < 0x20 || r == 0x7f }) >= 0 {
		return "/"
	}

	return u.String()
}

func LoginRedirect(next string) string {
	target := SafeNext(next)
	if target == "/" {
		return "/-/login"
	}

	return "/-/login?" + NextParam + "=" + url.QueryEscape(target)
}
