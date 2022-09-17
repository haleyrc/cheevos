package stringutil

import (
	"html"
	"strings"
)

func MakeSafe(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}
