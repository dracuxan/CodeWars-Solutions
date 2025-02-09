package kata

func BoolToWord(word bool) string {
	return map[bool]string{true: "Yes", false: "No"}[word]
}
