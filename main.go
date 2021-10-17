package main

import (
	"fmt"
	goenum "github.com/beloys/goenum/internal"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePathIn := os.Getenv("GOFILE")
	parser, err := goenum.New(filePathIn)
	if err != nil {
		panic(err)
	}
	ref, err := parser.Scan()
	if err != nil {
		panic(err)
	}
	factory, err := goenum.NewTypeFactory().Create(ref)
	if err != nil {
		panic(err)
	}
	w := strings.Builder{}
	err = goenum.NewTemplatePrinter(ref, factory).Print(&w)
	if err != nil {
		panic(err)
	}
	filePathOut := fmt.Sprintf("%s.generated.go", strings.Replace(filePathIn, ".go", "", 1))
	err = ioutil.WriteFile(filePathOut, []byte(w.String()), 0666)
	if err != nil {
		panic(err)
	}
}
