package main

import (
	"fmt"
	"github.com/geange/lucene-go/core/util/bytesutils"
)

func main() {
	pool := bytesutils.NewByteBlockPool(bytesutils.NewBytesAllocator(bytesutils.BYTE_BLOCK_SIZE, &bytesutils.DirectBytesAllocator{}))
	hash := bytesutils.NewBytesHash(pool)

	id, err := hash.Add([]byte("hello"))
	if err != nil {
		panic(err)
	}

	id2, err := hash.Add([]byte("hello world"))
	if err != nil {
		panic(err)
	}

	find := hash.Find([]byte("hello"))

	if id != find {
		panic(fmt.Sprintf("%d != %d", id, find))
	}

	fmt.Println(string(hash.Get(id)))
	fmt.Println(string(hash.Get(id2)))
}
