package main

import (
	"fmt"

	"github.com/geange/lucene-go/core/util/ints"
)

func main() {
	pool := ints.NewBlockPool(nil)

	writer := ints.NewSliceWriter(pool)

	writer.StartNewSlice()

	for i := 0; i < 100; i++ {
		writer.WriteInt(i)
	}

	writer.StartNewSlice()

	for i := 0; i < 100; i++ {
		writer.WriteInt(i)
	}

	fmt.Println(pool)
}
