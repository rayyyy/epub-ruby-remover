package internal

import (
	"regexp"
)

// RemoveRuby removes ruby tags from html.
func RemoveRuby(html string) string {
	// Remove rt tags and its content
	re := regexp.MustCompile(`<rt>.*?</rt>`)
	html = re.ReplaceAllString(html, "")

	// Remove ruby tag only
	re = regexp.MustCompile(`</?ruby>`)
	html = re.ReplaceAllString(html, "")

	return html
}
