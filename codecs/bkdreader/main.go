package main

import (
	"context"
	"github.com/geange/lucene-go/codecs/simpletext"
	coreIndex "github.com/geange/lucene-go/core/index"
	"github.com/geange/lucene-go/core/interface/index"
	"github.com/geange/lucene-go/core/store"
	"github.com/geange/lucene-go/core/util/version"
)

func main() {
	dir, err := store.NewNIOFSDirectory("data")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	segmentInfo := coreIndex.NewSegmentInfo(dir, version.Last, version.Last, "", 2048)

	state := index.NewSegmentWriteState(dir, segmentInfo, nil, nil, nil)

	writer, err := simpletext.NewPointsWriter(ctx, state)
	if err != nil {
		return
	}
	writer.WriteField()
}
