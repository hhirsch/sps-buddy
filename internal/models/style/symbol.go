package style

import (
	"regexp"
)

// mixed camel case also allows symbols to end in numbers
func IsMixedCamelCase(input string) bool {
	expression := regexp.MustCompile(`^[a-z]+([A-Z][a-z0-9]*)*$`)
	return expression.MatchString(input)
}

func IsCamelCase(input string) bool {
	expression := regexp.MustCompile(`^[a-z]+([A-Z][a-z]*)*$`)
	return expression.MatchString(input)
}
