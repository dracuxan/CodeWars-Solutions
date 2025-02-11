package kata

import "strings"

func GetCount(str string) (count int) {
	str_l := strings.Split(str, "")
	vovels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	count = 0
	for _, v := range str_l {
		if vovels[v] {
			count++
		}
	}

	return count
}
