package main

import (
	"context"
	"fmt"

	_ "github.com/geange/lucene-go/codecs/simpletext"
	"github.com/geange/lucene-go/core/index"
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
	topDocs, err := searcher.SearchTopN(context.Background(), search.NewMatchAllDocsQuery(), 100)
	if err != nil {
		panic(err)
	}

	result := topDocs.GetScoreDocs()
	for _, scoreDoc := range result {
		docID := scoreDoc.GetDoc()
		document, err := reader.Document(context.Background(), docID)
		if err != nil {
			panic(err)
		}

		for field := range document.GetFields("a") {
			fmt.Printf("段内排序后的文档号: %d  VS 段内排序前的文档: %s\n",
				scoreDoc.GetDoc(), field.Get())
		}
	}

	//searchSortField1 := index.NewSortedSetSortFieldV1("sort0", true, index.MAX)
	//searchSortField2 := index.NewSortedSetSortFieldV1("sort1", true, index.MIN)
	//searchSortFields := []index.SortField{searchSortField1, searchSortField2}
	//searchSort := index.NewSort(searchSortFields)
	////
	//search.

	//docs := reader.NumDocs()
	//
	//searcher.Search(search.NewNamedMatches(), 100, searchSort).coreDocs

	//fmt.Println(docs)
	//
	//{
	//	doc := document.NewDocument()
	//	doc.Add(document.NewStoredFieldAny("a", 74, document.STORED_ONLY))
	//	doc.Add(document.NewStoredFieldAny("a1", 86, document.STORED_ONLY))
	//	doc.Add(document.NewStoredFieldAny("a2", 1237, document.STORED_ONLY))
	//	docID, err := writer.AddDocument(doc)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(docID)
	//}
	//
	//{
	//	doc := document.NewDocument()
	//	doc.Add(document.NewStoredFieldAny("a", 123, document.STORED_ONLY))
	//	doc.Add(document.NewStoredFieldAny("a1", 123, document.STORED_ONLY))
	//	doc.Add(document.NewStoredFieldAny("a2", 789, document.STORED_ONLY))
	//
	//	docID, err := writer.AddDocument(doc)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(docID)
	//}
	//
	//{
	//	doc := document.NewDocument()
	//	doc.Add(document.NewStoredFieldAny("a", 741, document.STORED_ONLY))
	//	doc.Add(document.NewStoredFieldAny("a1", 861, document.STORED_ONLY))
	//	doc.Add(document.NewStoredFieldAny("a2", 12137, document.STORED_ONLY))
	//	docID, err := writer.AddDocument(doc)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(docID)
	//}
}
