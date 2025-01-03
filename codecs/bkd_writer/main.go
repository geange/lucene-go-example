package main

import (
	"context"
	"encoding/binary"
	"github.com/geange/lucene-go/codecs/simpletext"
	"github.com/geange/lucene-go/core/store"
	"github.com/geange/lucene-go/core/util/bkd"
)

func main() {
	dir, err := store.NewNIOFSDirectory("data")
	if err != nil {
		panic(err)
	}

	cfg, err := bkd.NewConfig(2, 2, 4, 2)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	output, err := dir.CreateOutput(ctx, "bkd.txt")
	if err != nil {
		panic(err)
	}

	writer := simpletext.NewBKDWriter(
		100, dir, "demo", cfg, 16, 4)

	err = writer.Add(ctx, Point(5, 4), 1)
	if err != nil {
		return
	}
	err = writer.Add(ctx, Point(1, 2), 1)
	if err != nil {
		return
	}
	err = writer.Add(ctx, Point(1, 3), 1)
	if err != nil {
		return
	}
	err = writer.Add(ctx, Point(2, 9), 2)
	if err != nil {
		return
	}

	writer.Finish(output)
	output.Close()
}

func Point(values ...int) []byte {
	size := 4 * len(values)
	bs := make([]byte, size)
	for i := 0; i < len(values); i++ {
		binary.BigEndian.PutUint32(bs[i*4:], uint32(values[i]))
	}
	return bs
}
