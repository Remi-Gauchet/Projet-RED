package piscine

func Capitalize(s string) string {
	var chainerune []rune = []rune(s)
	var nouveaumot bool = true
	for indexrune := 0; indexrune < len(chainerune); indexrune++ {
		c := chainerune[indexrune]
		alphanum := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
		if alphanum {
			if nouveaumot == true {
				if c >= 'a' && c <= 'z' {
					chainerune[indexrune] = c - ('a' - 'A')
				}
				nouveaumot = false
			} else if c >= 'A' && c <= 'Z' {
				chainerune[indexrune] = c + ('a' - 'A')
			}
		} else {
			nouveaumot = true
		}
	}
	return string(chainerune)
}
