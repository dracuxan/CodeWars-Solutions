package kata

func duplicate_count(s1 string) int {
	counter := 0
	l := make(map[rune]bool)

	for _, char := range s1 {
		if l[char] {
			counter++
			continue
		} else {
			l[char] = true
		}
	}

	return 0
}
