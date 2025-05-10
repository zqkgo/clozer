package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zqkgo/clozer"
)

var (
	path   string
	symbol string

	defaultPath = "./content"
	fs          flag.FlagSet
)

func init() {
	fs.StringVar(&path, "path", "", "path of file containing content")
	fs.StringVar(&symbol, "symbol", "x", "symbol to cloze")
}

func main() {
	// clozer type
	t := os.Args[1]
	fs.Parse(os.Args[2:])
	log.Println("args: ", os.Args)
	if path == "" {
		path = defaultPath
		log.Printf("path not be specified, will use default path: %s\n", path)
	}
	f, err := os.OpenFile(path, os.O_RDONLY, 0766)
	if err != nil {
		log.Fatalf("failed to open file: %s, error: %+v\n", path, err)
	}
	c := clozer.GetClozer(t)
	s, err := c.Cloze(f, clozer.WithSymbol(symbol))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
