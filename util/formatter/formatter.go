package formatter

import "regexp"

func RemoveMaskFromDocument(document string) string {
	strRegexp := regexp.MustCompile(`[./-]`)
	unmaskedDocument := strRegexp.ReplaceAllString(document, "")
	return unmaskedDocument
}
