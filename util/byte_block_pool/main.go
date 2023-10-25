package main

import (
	"fmt"
	"github.com/geange/lucene-go/core/util/bytesutils"
)

func main() {
	pool := bytesutils.NewByteBlockPool(bytesutils.NewBytesAllocator(bytesutils.BYTE_BLOCK_SIZE, &bytesutils.DirectBytesAllocator{}))
	pool.NewSlice(2)
	pool.Append([]byte("abcdefg"))

	pool.Append([]byte("abcdefg"))

	fmt.Println(pool)
}
