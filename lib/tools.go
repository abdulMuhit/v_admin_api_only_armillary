package lib

import (
	"regexp"
	"strings"
)

//Slug is lib for create slug
func Slug(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]")
	var slug = reg.ReplaceAllString(str, "-")
	reg, _ = regexp.Compile("[-]+")
	slug = reg.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return strings.ToLower(slug)
}
