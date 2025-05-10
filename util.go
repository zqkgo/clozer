package clozer

func MultiByteRune(r rune) bool {
	letter := r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
	return !letter
}

func IsLetter(r byte) bool {
	return r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
}
