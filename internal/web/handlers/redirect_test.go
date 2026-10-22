package handlers

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSafeNext(t *testing.T) {
	tests := []struct {
		name   string
		target string
		want   string
	}{
		{"empty", "", "/"},
		{"root", "/", "/"},
		{"path", "/thomas/abc", "/thomas/abc"},
		{"path with query", "/thomas/abc?file=main.go&line=2", "/thomas/abc?file=main.go&line=2"},
		{"path with fragment", "/thomas/abc#file", "/thomas/abc#file"},
		{"query containing a URL", "/thomas/abc?next=https://evil.example", "/thomas/abc?next=https://evil.example"},
		{"relative path", "thomas/abc", "/"},
		{"absolute URL", "https://evil.example/thomas/abc", "/"},
		{"protocol relative URL", "//evil.example/thomas/abc", "/"},
		{"userinfo", "//user@evil.example/thomas/abc", "/"},
		{"backslash", "/\\evil.example/thomas/abc", "/"},
		{"encoded backslash", "/%5Cevil.example/thomas/abc", "/"},
		{"encoded protocol relative URL", "/%2F%2Fevil.example/thomas/abc", "/"},
		{"header injection", "/foo%0D%0ALocation:%20https://evil.example", "/"},
		{"control character", "/foo%09bar", "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, SafeNext(tt.target))
		})
	}
}

func TestLoginRedirect(t *testing.T) {
	require.Equal(t, "/-/login", LoginRedirect(""))
	require.Equal(t, "/-/login", LoginRedirect("https://evil.example"))
	require.Equal(t, "/-/login?next=%2Fthomas%2Fabc%3Ffile%3Dmain.go", LoginRedirect("/thomas/abc?file=main.go"))
	require.Equal(t, "/-/login?next="+url.QueryEscape("/thomas/abc"), LoginRedirect("/thomas/abc"))
}
