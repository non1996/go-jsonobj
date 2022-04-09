package wrapper

type String struct {
	s     string
	valid bool
}

func (s *String) ParseString(str string) (err error) {
	s.s = str
	s.valid = true
	return nil
}

func (s *String) ValueOf(v string) {
	s.s = v
	s.valid = true
}

func (s *String) Value() string {
	return s.V()
}

func (s *String) V() string {
	if s.valid {
		return s.s
	}
	return ""
}

func (s *String) Valid() bool {
	return s.valid
}

func (s *String) IsEmpty() bool {
	return len(s.s) == 0
}
