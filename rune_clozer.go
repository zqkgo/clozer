package main

import (
	"bytes"
	"fmt"
	"io"
)

type runeClozer struct{}

func (tc *runeClozer) Cloze(rd io.ReadCloser) (string, error) {
	bs, err := io.ReadAll(rd)
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
		if !ignRunes[c] {
			if idx%2 == 0 {
				s1 = replaceChar(s1, true)
			} else {
				s2 = replaceChar(s1, true)
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

func replaceChar(c string, cloze bool) string {
	if !cloze {
		return symbol
	}
	return "{{c1::" + c + "}}"
}
