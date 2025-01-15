package main

import (
	"context"
	"fmt"

	_ "github.com/geange/lucene-go/codecs/simpletext"
	"github.com/geange/lucene-go/core/index"
	index2 "github.com/geange/lucene-go/core/interface/index"
	"github.com/geange/lucene-go/core/search"
	"github.com/geange/lucene-go/core/store"
)

func main() {
	dir, err := store.NewNIOFSDirectory("data")
	if err != nil {
		panic(err)
	}

	reader, err := index.OpenDirectoryReader(context.Background(), dir, nil, nil)
	if err != nil {
		panic(err)
	}

	searcher, err := search.NewIndexSearcher(reader)
	if err != nil {
		panic(err)
	}

	q1 := search.NewTermQuery(index.NewTerm("content", []byte("a")))
	q2 := search.NewTermQuery(index.NewTerm("content", []byte("c")))
	q3 := search.NewTermQuery(index.NewTerm("content", []byte("e")))
	q4 := search.NewTermQuery(index.NewTerm("author", []byte("author4")))

	builder := search.NewBooleanQueryBuilder()
	builder.AddQuery(q1, index2.OccurMust)
	builder.AddQuery(q2, index2.OccurMust)
	builder.AddQuery(q3, index2.OccurMust)
	builder.AddQuery(q4, index2.OccurMust)
	query, err := builder.Build()
	if err != nil {
		panic(err)
	}

	topDocs, err := searcher.SearchTopN(context.Background(), query, 5)
	if err != nil {
		panic(err)
	}

	for i, doc := range topDocs.GetScoreDocs() {
		fmt.Printf("result%d: 文档%d\n", i, doc.GetDoc())
	}
}
