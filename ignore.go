package main

func init() {
	for _, word := range []string{
		"with",
		"and",
		"or",
		"the",
		"as",
		"well",
		"good",
		"bad",

		"to",
		"via",
		"at",
		"on",

		"you",
		"he",
		"she",
		"it",
	} {
		ignWords[word] = true
	}
}

var ignRunes map[rune]bool = map[rune]bool{
	'，': true,
	'。': true,
	'！': true,
	'？': true,
	'；': true,
	'“': true,
	'”': true,
	'、': true,
}

var ignWords map[string]bool = map[string]bool{}
