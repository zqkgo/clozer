package clozer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type codeClozer struct{}

func (cc *codeClozer) Cloze(rc io.ReadCloser, opts ...ClozeOpt) (string, error) {
	defer rc.Close()
	var bf bytes.Buffer
	s := bufio.NewScanner(rc)
	idx := 1
	for s.Scan() {
		t := s.Text()
		// t := strings.TrimSpace(orit)
		if !cc.shouldSkip(t) {
			pure := strings.TrimSpace(t)
			ct := fmt.Sprintf("{{c%d::%s}}\n", idx, pure)
			idx++
			t = strings.Replace(t, pure, ct, 1)
		} else {
			t = t + "\n"
		}
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
