package clozer

import "io"

const (
	typeCodeClozer = "code"
	typeRuneClozer = "rune"
	typeWordClozer = "word"
)

type ClozeOpt func(o *ClozeOptions)

func WithToggle(tgl bool) ClozeOpt {
	return func(o *ClozeOptions) {
		o.Toggle = tgl
	}
}

func WithSymbol(s string) ClozeOpt {
	return func(o *ClozeOptions) {
		o.Symbol = s
	}
}

type ClozeOptions struct {
	// 交替 cloze。
	Toggle bool
	Symbol string
}

type Clozer interface {
	Cloze(io.ReadCloser, ...ClozeOpt) (string, error)
}

func GetClozer(t string) Clozer {
	switch t {
	case typeCodeClozer:
		return &codeClozer{}
	case typeRuneClozer:
		return &runeClozer{}
	case typeWordClozer:
		return &wordClozer{}
	default:
		return &codeClozer{}
	}
}
