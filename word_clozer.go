package clozer

import (
	"errors"
	"io"
)

type wordClozer struct{}

func (cc *wordClozer) Cloze(rc io.ReadCloser, opts ...ClozeOpt) (string, error) {
	defer rc.Close()
	bs, err := io.ReadAll(rc)
	if err != nil {
		return "", err
	}
	if len(bs) == 0 {
		return "", errors.New("empty text")
	}
	r1 := clozeWord(bs, false)
	r2 := clozeWord(bs, true)
	return r1 + "\n" + r2, nil
}

func clozeWord(bs []byte, toggleCloze bool) string {
	var idx int
	var result []byte
	for idx < len(bs) {
		// 跳过开头连续的空格。
		if len(result) == 0 && bs[idx] == ' ' {
			idx++
			continue
		}
		// 非字母原样放回。
		for idx < len(bs) {
			if !IsLetter(bs[idx]) {
				result = append(result, bs[idx])
				idx++
				continue
			}
			break
		}
		// 找到下一个单词。
		var word []byte
		for idx < len(bs) {
			if !IsLetter(bs[idx]) {
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
	return string(result)
}
