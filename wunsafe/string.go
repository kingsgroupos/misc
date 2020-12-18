package wunsafe

import "strconv"

type String []byte

func (s String) Equals(str string) bool {
	if len(s) > 0 {
		return BytesToString(s) == str
	} else {
		return str == ""
	}
}

func (s String) Atoi64() (int64, error) {
	const fnAtoi64 = "Atoi64"

	sLen := len(s)
	if 0 < sLen && sLen < 19 {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, &strconv.NumError{fnAtoi64, string(s0), strconv.ErrSyntax}
			}
		}

		n := int64(0)
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0, &strconv.NumError{fnAtoi64, string(s0), strconv.ErrSyntax}
			}
			n = n*10 + int64(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n, nil
	}

	// Slow path for invalid, big, or underscored integers.
	i64, err := strconv.ParseInt(string(s), 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnAtoi64
	}
	return i64, err
}

func (s String) Atou64() (uint64, error) {
	const fnAtou64 = "Atou64"

	sLen := len(s)
	if 0 < sLen && sLen < 19 {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' {
			return 0, &strconv.NumError{fnAtou64, string(s0), strconv.ErrSyntax}
		}
		if s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, &strconv.NumError{fnAtou64, string(s0), strconv.ErrSyntax}
			}
		}

		n := uint64(0)
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0, &strconv.NumError{fnAtou64, string(s0), strconv.ErrSyntax}
			}
			n = n*10 + uint64(ch)
		}
		return n, nil
	}

	// Slow path for invalid, big, or underscored integers.
	i64, err := strconv.ParseUint(string(s), 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnAtou64
	}
	return i64, err
}

func (s String) String() string {
	return string(s)
}

func (s String) UnsafeString() string {
	if len(s) > 0 {
		return BytesToString(s)
	} else {
		return ""
	}
}
