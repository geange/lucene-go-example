package main

import (
	"context"
	"fmt"
	"os"

	"github.com/geange/lucene-go/codecs/simpletext"
	"github.com/geange/lucene-go/core/document"
	"github.com/geange/lucene-go/core/index"
	"github.com/geange/lucene-go/core/search"
	"github.com/geange/lucene-go/core/store"
)

func main() {
	err := os.RemoveAll("data")
	if err != nil {
		panic(err)
	}

	dir, err := store.NewNIOFSDirectory("data")
	if err != nil {
		panic(err)
	}

	codec := simpletext.NewCodec()
	similarity, err := search.NewBM25Similarity()
	if err != nil {
		panic(err)
	}

	config := index.NewIndexWriterConfig(codec, similarity)

	writer, err := index.NewIndexWriter(context.Background(), dir, config)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	{
		doc := document.NewDocument()
		doc.Add(document.NewTextField("a", "74", true))
		doc.Add(document.NewTextField("a1", "86", true))
		doc.Add(document.NewTextField("a2", "1237", true))
		docID, err := writer.AddDocument(context.Background(), doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(docID)
	}

	{
		doc := document.NewDocument()
		doc.Add(document.NewTextField("a", "123", true))
		doc.Add(document.NewTextField("a1", "123", true))
		doc.Add(document.NewTextField("a2", "789", true))

		docID, err := writer.AddDocument(context.Background(), doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(docID)
	}

	{
		doc := document.NewDocument()
		doc.Add(document.NewTextField("a", "741", true))
		doc.Add(document.NewTextField("a1", "861", true))
		doc.Add(document.NewTextField("a2", "12137", true))
		docID, err := writer.AddDocument(context.Background(), doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(docID)
	}

	//{
	//	doc := document.NewDocument()
	//	point1, _ := document.NewBinaryPoint("p1", []byte{1}, []byte{2})
	//	point2, _ := document.NewBinaryPoint("p2", []byte{1}, []byte{2})
	//	point3, _ := document.NewBinaryPoint("p3", []byte{1}, []byte{2})
	//	doc.Add(point1)
	//	doc.Add(point2)
	//	doc.Add(point3)
	//	docID, err := writer.AddDocument(context.Background(), doc)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(docID)
	//}

}
