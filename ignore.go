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

	for _, r := range []rune{
		'，',
		'。',
		'！',
		'？',
		'；',
		'：',
		'“',
		'”',
		'、',
		'（',
		'）',
		'(',
		')',
		' ',
		'.',
		':',
	} {
		ignRunes[r] = true
	}
}

var ignRunes map[rune]bool = map[rune]bool{}
var ignWords map[string]bool = map[string]bool{}
