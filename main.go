package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	typeCodeClozer = "code"
	typeRuneClozer = "rune"
	typeWordClozer = "word"
)

var (
	path   string
	symbol string

	defaultPath = "./content"
	fs          flag.FlagSet
)

type Clozer interface {
	Cloze(io.ReadCloser) (string, error)
}

func init() {
	fs.StringVar(&path, "path", "", "path of file containing content")
	fs.StringVar(&symbol, "symbol", "x", "symbol to cloze")
}

func main() {
	// clozer type
	t := os.Args[1]
	fs.Parse(os.Args[2:])
	if path == "" {
		path = defaultPath
		log.Printf("path not be specified, will use default path: %s\n", path)
	}
	f, err := os.OpenFile(path, os.O_RDONLY, 0766)
	if err != nil {
		log.Fatalf("failed to open file: %s, error: %+v\n", path, err)
	}
	c := getClozer(t)
	s, err := c.Cloze(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
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
