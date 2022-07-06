package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
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

func (tc *textClozer) Cloze(rc io.ReadCloser) (string, error) {
	bs, err := ioutil.ReadAll(rc)
	if err != nil {
		return "", err
	}
	defer rc.Close()
	txt := []rune(string(bs))
	// open and write cloze files
	p1, p2 := path+".cloze.1", path+".cloze.2"
	clozeFile, err := os.OpenFile(p1, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalf("failed to create cloze file: %s, error: %+v\n", p1, err)
	}
	defer clozeFile.Close()
	clozeRevFile, err := os.OpenFile(p2, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalf("failed to create cloze reverse file: %s, error: %+v\n", p2, err)
	}
	defer clozeRevFile.Close()

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
		_, err = clozeFile.WriteString(s1)
		if err != nil {
			log.Fatalf("failed to write cloze file: %s, error: %+v\n", p1, err)
		}
		_, err = clozeRevFile.WriteString(s2)
		if err != nil {
			log.Fatalf("failed to write cloze file: %s, error: %+v\n", p2, err)
		}
	}
	return "", nil
}