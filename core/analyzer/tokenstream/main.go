package main

import (
	"fmt"
	"github.com/geange/lucene-go/core/analysis"
	"github.com/geange/lucene-go/core/analysis/standard"
)

func main() {
	set := analysis.NewCharArraySet()
	set.Add(" ")
	set.Add("\n")
	set.Add("\t")

	analyzer := standard.NewAnalyzer(set)

	imp := analysis.NewAnalyzer(analyzer)

	stream, err := imp.GetTokenStreamFromText("xxxx", "aaaa BBBFFDs cccc dddd")
	if err != nil {
		panic(err)
	}

	_, err = stream.IncrementToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(stream.AttributeSource().CharTerm().Buffer()))

	_, err = stream.IncrementToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(stream.AttributeSource().CharTerm().Buffer()))

	_, err = stream.IncrementToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(stream.AttributeSource().CharTerm().Buffer()))
}
