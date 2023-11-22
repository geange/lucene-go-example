package main

import (
	"fmt"
	"github.com/geange/lucene-go/core/document"

	"github.com/geange/lucene-go/codecs/simpletext"
	"github.com/geange/lucene-go/core/index"
	"github.com/geange/lucene-go/core/search"
	"github.com/geange/lucene-go/core/store"
)

func main() {
	dir, err := store.NewNIOFSDirectory("data")
	if err != nil {
		panic(err)
	}

	codec := simpletext.NewCodec()
	similarity := search.NewCastBM25Similarity()

	config := index.NewWriterConfig(codec, similarity)

	writer, err := index.NewWriter(dir, config)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	{
		doc := document.NewDocument()
		doc.Add(document.NewField[int32]("a", 74, document.STORED_ONLY))
		doc.Add(document.NewField[int32]("a1", 86, document.STORED_ONLY))
		doc.Add(document.NewField[int32]("a2", 1237, document.STORED_ONLY))
		docID, err := writer.AddDocument(doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(docID)
	}

	{
		doc := document.NewDocument()
		doc.Add(document.NewField[int32]("a", 123, document.STORED_ONLY))
		doc.Add(document.NewField[int32]("a1", 123, document.STORED_ONLY))
		doc.Add(document.NewField[int32]("a2", 789, document.STORED_ONLY))

		docID, err := writer.AddDocument(doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(docID)
	}

	{
		doc := document.NewDocument()
		doc.Add(document.NewField[int32]("a", 741, document.STORED_ONLY))
		doc.Add(document.NewField[int32]("a1", 861, document.STORED_ONLY))
		doc.Add(document.NewField[int32]("a2", 12137, document.STORED_ONLY))
		docID, err := writer.AddDocument(doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(docID)
	}

	{
		doc := document.NewDocument()
		point1, _ := document.NewBinaryPoint("p1", []byte{1}, []byte{2})
		point2, _ := document.NewBinaryPoint("p2", []byte{1}, []byte{2})
		point3, _ := document.NewBinaryPoint("p3", []byte{1}, []byte{2})
		doc.Add(point1)
		doc.Add(point2)
		doc.Add(point3)
		docID, err := writer.AddDocument(doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(docID)
	}

	{

	}

}
