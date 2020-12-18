package misc

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"
	"unicode"
)

func Split(str string, sep string) []string {
	var r []string
	a := strings.Split(str, sep)
	for _, s := range a {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) > 0 {
			r = append(r, trimmed)
		}
	}
	return r
}

func UnescapeUnicode(str string) string {
	if !strings.Contains(str, `\u`) {
		return str
	}

	var out bytes.Buffer
	var i int
	for i <= len(str)-6 {
		if str[i] == '\\' && str[i+1] == 'u' {
			if u, err := strconv.ParseUint(str[i+2:i+6], 16, 64); err == nil {
				out.WriteRune(rune(u))
				i += 6
				continue
			}
		}
		out.WriteByte(str[i])
		i++
	}

	if i < len(str) {
		out.WriteString(str[i:])
	}
	return out.String()
}

func UCFirst(str string) string {
	var r string
	for i, v := range str {
		switch i {
		case 0:
			r = string(unicode.ToUpper(v))
		default:
			return r + str[i:]
		}
	}
	return r
}

func LCFirst(str string) string {
	var r string
	for i, v := range str {
		switch i {
		case 0:
			r = string(unicode.ToLower(v))
		default:
			return r + str[i:]
		}
	}
	return r
}

func IndexStrings(a []string, str string) int {
	for i, s := range a {
		if s == str {
			return i
		}
	}
	return -1
}

func DecodeCompileTimeString(str string) (string, error) {
	str = strings.ReplaceAll(str, "#", "=")
	bts, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimRight(bts, "\r\n")), nil
}
