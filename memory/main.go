package main

import (
	"fmt"

	"github.com/geange/lucene-go/core/analysis"
	"github.com/geange/lucene-go/core/analysis/standard"
	"github.com/geange/lucene-go/core/document"
	"github.com/geange/lucene-go/core/index"
	"github.com/geange/lucene-go/memory"
)

func main() {
	//mi, err := memory.NewIndex()
	//if err != nil {
	//	panic(err)
	//}
	//
	//set := analysis.NewCharArraySet()
	//set.Add(" ")
	//set.Add("\n")
	//set.Add("\t")
	//
	//analyzer := standard.NewAnalyzer(set)
	//
	//field := document.NewTextField("f1", "some some some some some some some vv", false)
	//err = mi.AddIndexAbleField(field, analyzer)
	//if err != nil {
	//	panic(err)
	//}
	//mi.Freeze()
	//
	//count := mi.Search(search.NewTermQuery(index.NewTerm("f1", []byte("some"))))
	//fmt.Println(count)
	testPointValuesDoNotAffectPositionsOrOffset()
}

func testPointValuesDoNotAffectPositionsOrOffset() error {
	mi, err := memory.NewIndex(
		memory.WithStorePayloads(true),
		memory.WithStoreOffsets(true),
	)
	if err != nil {
		return err
	}

	set := analysis.NewCharArraySet()
	set.Add(" ")
	set.Add("\n")
	set.Add("\t")
	analyzer := standard.NewAnalyzer(set)

	err = mi.AddIndexAbleField(document.NewTextField("text", "quick brown fox", false), analyzer)
	if err != nil {
		return err
	}
	point1, err := document.NewBinaryPoint("text", []byte("quick"))
	if err != nil {
		return err
	}
	err = mi.AddIndexAbleField(point1, analyzer)
	if err != nil {
		return err
	}

	point2, err := document.NewBinaryPoint("text", []byte("brown"))
	if err != nil {
		return err
	}
	err = mi.AddIndexAbleField(point2, analyzer)
	if err != nil {
		return err
	}

	leaves, err := mi.CreateSearcher().GetIndexReader().Leaves()
	if err != nil {
		return err
	}
	leafReader := leaves[0].Reader()
	terms, err := leafReader.(index.LeafReader).Terms("text")
	if err != nil {
		return err
	}
	tenum, err := terms.Iterator()
	if err != nil {
		return err
	}
	{
		bytes, err := tenum.Next()
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
	}

	{
		bytes, err := tenum.Next()
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
	}

	{
		bytes, err := tenum.Next()
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
	}
	return nil
}
