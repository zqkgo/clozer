package main

import "io"

type clozeOpt func(o *clozeOptions)

func withToggle(tgl bool) clozeOpt {
	return func(o *clozeOptions) {
		o.toggle = tgl
	}
}

type clozeOptions struct {
	// 交替 cloze。
	toggle bool
}

type Clozer interface {
	Cloze(io.ReadCloser, ...clozeOpt) (string, error)
}

func getClozer(t string) Clozer {
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
