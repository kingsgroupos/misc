package wlang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	rexPos    = regexp.MustCompile(`{\d+[^}]*}`)
	rexNumFmt = regexp.MustCompile(`^:[0#]*\.([0#]+)`)
)

func Sprint(lang string, a ...interface{}) (string, error) {
	l, err := language.Parse(lang)
	return message.NewPrinter(l).Sprint(a...), err
}

func Sprintln(lang string, a ...interface{}) (string, error) {
	l, err := language.Parse(lang)
	return message.NewPrinter(l).Sprintln(a...), err
}

func convertFormat(format string) string {
	if strings.Contains(format, "{{") {
		format = strings.Replace(format, "{{", "{", -1)
	}
	if strings.Contains(format, "}}") {
		format = strings.Replace(format, "}}", "}", -1)
	}
	if strings.Contains(format, "%") {
		format = strings.Replace(format, "%", "%%", -1)
	}
	return rexPos.ReplaceAllStringFunc(format, func(s string) string {
		i := 2
		for n := len(s); i < n; i++ {
			if ch := s[i]; ch < '0' || ch > '9' {
				break
			}
		}
		v, _ := strconv.Atoi(s[1:i])
		if s[i] == ':' {
			a := rexNumFmt.FindStringSubmatch(s[i:])
			if len(a) == 2 {
				return fmt.Sprintf("%%.%d[%d]f", len(a[1]), v+1)
			}
		}
		return fmt.Sprintf("%%[%d]v", v+1)
	})
}

func Sprintf(lang string, format string, a ...interface{}) (string, error) {
	l, err := language.Parse(lang)
	f := convertFormat(format)
	return message.NewPrinter(l).Sprintf(f, a...), err
}
