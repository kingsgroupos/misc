package variable

import (
	"strings"
	"unicode"

	"github.com/kingsgroupos/misc"
)

const (
	validVariableChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
)

func splitVariableName(name string) []string {
	var sb strings.Builder
	sb.Grow(len(name))
	var prev = '_'
	for _, r := range name {
		if strings.IndexRune(validVariableChars, r) >= 0 {
			if prev == '_' && r == '_' {
				continue
			}
			if sb.Len() == 0 && unicode.IsDigit(r) {
				continue
			}
			if prev != '_' && r != '_' && unicode.IsDigit(r) != unicode.IsDigit(prev) {
				sb.WriteRune('_')
			}
			sb.WriteRune(r)
			prev = r
		} else {
			if prev == '_' {
				continue
			}
			sb.WriteByte('_')
			prev = '_'
		}
	}

	return misc.Split(sb.String(), "_")
}

func concatFields(fields []string) string {
	n := len(fields)
	if n > 1 && len(fields[n-1]) > 2 {
		idx := strings.IndexFunc(fields[n-1], func(r rune) bool {
			return r < '0' || r > '9'
		})
		if idx < 0 {
			return strings.Join(fields[:n-1], "") + "_" + fields[n-1]
		}
	}

	return strings.Join(fields, "")
}

func ToPascal(name string) string {
	fields := splitVariableName(name)
	if len(fields) == 0 {
		return ""
	}
	if c := fields[0][0]; c >= '0' && c <= '9' {
		return ""
	}
	for i, f := range fields {
		fields[i] = misc.UCFirst(f)
	}

	return concatFields(fields)
}

func ToCamel(name string) string {
	fields := splitVariableName(name)
	if len(fields) == 0 {
		return ""
	}
	for i, f := range fields {
		switch i {
		case 0:
			if c := f[0]; c >= '0' && c <= '9' {
				return ""
			}
			idx := strings.IndexFunc(f, func(r rune) bool {
				return r < 'A' || r > 'Z'
			})
			switch idx {
			case -1:
				fields[0] = strings.ToLower(f)
			case 0:
			case 1:
				fields[0] = strings.ToLower(f[:1]) + f[1:]
			default:
				if c := f[idx]; c >= 'a' && c <= 'z' {
					idx--
				}
				fields[0] = strings.ToLower(f[:idx]) + f[idx:]
			}
		default:
			fields[i] = misc.UCFirst(f)
		}
	}

	return concatFields(fields)
}

func ToProtoStyle(name string) string {
	isASCIILower := func(c byte) bool {
		return 'a' <= c && c <= 'z'
	}
	isASCIIDigit := func(c byte) bool {
		return '0' <= c && c <= '9'
	}

	if name == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if name[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(name); i++ {
		c := name[i]
		if c == '_' && i+1 < len(name) && isASCIILower(name[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(name) && isASCIILower(name[i+1]) {
			i++
			t = append(t, name[i])
		}
	}
	return string(t)
}

var keywords = map[string][]string{
	"go": {
		"break",
		"default",
		"func",
		"interface",
		"select",
		"case",
		"defer",
		"go",
		"map",
		"struct",
		"chan",
		"else",
		"goto",
		"package",
		"switch",
		"const",
		"fallthrough",
		"if",
		"range",
		"type",
		"continue",
		"for",
		"import",
		"return",
		"var",
	},
}

func FixConflicts(name string, langs ...string) string {
	if len(langs) == 0 {
		langs = []string{"go"}
	}
	for _, lang := range langs {
		a := keywords[lang]
		if a == nil {
			panic("unsupported language: " + lang)
		}
		if misc.IndexStrings(a, name) >= 0 {
			return name + "_"
		}
	}
	return name
}
