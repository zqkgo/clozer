package main

import (
	"errors"
	"io"
)

type wordClozer struct{}

func (cc *wordClozer) Cloze(rc io.ReadCloser) (string, error) {
	defer rc.Close()
	bs, err := io.ReadAll(rc)
	if err != nil {
		return "", err
	}
	if len(bs) == 0 {
		return "", errors.New("empty text")
	}

	var idx int
	var result []byte
	var toggleCloze bool
	for idx < len(bs) {
		// 跳过开头连续的空格。
		if len(result) == 0 && bs[idx] == ' ' {
			idx++
			continue
		}
		// 非字母原样放回。
		for idx < len(bs) {
			if !isLetter(bs[idx]) {
				result = append(result, bs[idx])
				idx++
				continue
			}
			break
		}
		// 找到下一个单词。
		var word []byte
		for idx < len(bs) {
			if !isLetter(bs[idx]) {
				break
			}
			word = append(word, bs[idx])
			idx++
		}
		// 遍历到最后。
		if len(word) == 0 {
			break
		}
		if toggleCloze {
			word = append([]byte("{{c1::"), word...)
			word = append(word, []byte("}}")...)
		}
		result = append(result, word...)
		toggleCloze = !toggleCloze
	}
	return string(result), nil
}

func isLetter(r byte) bool {
	return r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
}
