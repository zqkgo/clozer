package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type codeClozer struct{}

func (cc *codeClozer) Cloze(rc io.ReadCloser) (string, error) {
	defer rc.Close()
	var bf bytes.Buffer
	s := bufio.NewScanner(rc)
	idx := 1
	for s.Scan() {
		t := s.Text()
		// t := strings.TrimSpace(orit)
		if !cc.shouldSkip(t) {
			pure := t
			commentSign := strings.Index(pure, "//")
			if commentSign != -1 {
				pure = pure[:commentSign]
			}
			pure = strings.TrimSpace(pure)

			ct := fmt.Sprintf("{{c%d::%s}}", idx, pure)
			idx++
			t = strings.Replace(t, pure, ct, 1)

		}
		t += "\n"
		_, err := bf.WriteString(t)
		if err != nil {
			return "", err
		}
	}
	return bf.String(), nil
}

func (cc *codeClozer) shouldSkip(s string) bool {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "//") {
		return true
	}
	if s == "}" {
		return true
	}
	if len(s) == 0 {
		return true
	}
	return false
}
