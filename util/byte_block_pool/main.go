package main

import (
	"fmt"

	"github.com/geange/lucene-go/core/util/bytes"
)

func main() {
	pool := bytes.NewByteBlockPool(bytes.NewDirectAllocator(bytes.BYTE_BLOCK_SIZE))
	pool.NewSlice(2)
	pool.Append([]byte("abcdefg"))

	pool.Append([]byte("abcdefg"))

	fmt.Println(pool)
}
