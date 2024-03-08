package main

import (
	"bytes"
	"fmt"
	"io"
)

type runeClozer struct{}

func (tc *runeClozer) Cloze(rd io.ReadCloser, opts ...clozeOpt) (string, error) {
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

		// 非 rune 的字符，例如英语单词。
		var notRuneWord []rune
		for i < len(txt) {
			if !isRuneLetter(txt[i]) {
				break
			}
			notRuneWord = append(notRuneWord, txt[i])
			i++
		}
		if len(notRuneWord) > 0 {
			bf1.WriteString(string(notRuneWord))
			bf2.WriteString(string(notRuneWord))
			continue
		}

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
