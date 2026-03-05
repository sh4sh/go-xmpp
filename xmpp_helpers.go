package xmpp

import (
	"strings"
)

func Unquote(msg string) string {
	var b strings.Builder
	b.Grow(len(msg))

	lines := strings.SplitAfter(msg, "\n")
	for _, line := range lines {
		b.WriteString(strings.TrimPrefix(line, "> "))
	}
	return b.String()
}

// Quote prefixes each line with "> " per XEP-0393, producing the fallback
// block that accompanies a XEP-0461 reply. It is the inverse of Unquote.
func Quote(msg string) string {
	var b strings.Builder
	b.Grow(len(msg) + strings.Count(msg, "\n")*2 + 2)

	lines := strings.SplitAfter(msg, "\n")
	for _, line := range lines {
		if line != "" {
			b.WriteString("> ")
			b.WriteString(line)
		}
	}
	return b.String()
}
