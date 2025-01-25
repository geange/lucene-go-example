package main

import (
	"fmt"

	"github.com/geange/lucene-go/core/util/bytesref"
)

func main() {
	pool := bytesref.NewBlockPool(bytesref.GetAllocatorBuilder().NewDirect(100))
	pool.NewSlice(2)
	pool.Append([]byte("abcdefg"))

	pool.Append([]byte("abcdefg"))

	fmt.Println(pool.Get(0))
}
