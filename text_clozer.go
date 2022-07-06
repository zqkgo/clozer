package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

var ignChars map[rune]bool = map[rune]bool{
	'，': true,
	'。': true,
	'！': true,
	'？': true,
	'；': true,
	'“': true,
	'”': true,
	'、': true,
}

type textClozer struct{}

func (tc *textClozer) Cloze(rd io.ReadCloser) (string, error) {
	bs, err := ioutil.ReadAll(rd)
	if err != nil {
		return "", err
	}
	defer rd.Close()
	txt := []rune(string(bs))
	var bf1, bf2 bytes.Buffer

	idx := 0
	for i := 0; i < len(txt); i++ {
		c := txt[i]
		s1, s2 := string(c), string(c)
		if !ignChars[c] {
			if idx%2 == 0 {
				s1 = symbol
			} else {
				s2 = symbol
			}
			idx++
		}
		_, err = bf1.WriteString(s1)
		if err != nil {
			return "", fmt.Errorf("failed to write buffer, error: %+v", err)
		}
		_, err = bf2.WriteString(s2)
		if err != nil {
			return "", fmt.Errorf("failed to write buffer, error: %+v", err)
		}
	}
	return bf1.String() + "\n" + bf2.String(), nil
}
