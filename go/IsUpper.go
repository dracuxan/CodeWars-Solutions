package kata

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return false
		}
	}
	return true
}
