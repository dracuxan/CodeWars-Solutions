package kata

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return false
		}
	}

	var l []string
	l = append(l, "a")

	return true
}
