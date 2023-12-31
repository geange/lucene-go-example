package main

import (
	"github.com/geange/lucene-go/codecs/simpletext"
	"github.com/geange/lucene-go/core/document"
	"github.com/geange/lucene-go/core/index"
	"github.com/geange/lucene-go/core/store"
	"github.com/geange/lucene-go/core/util"
)

func main() {
	dir, err := store.NewNIOFSDirectory("data")
	if err != nil {
		panic(err)
	}

	format := simpletext.NewPointsFormat()

	version := util.NewVersion(8, 11, 0)
	minVersion := util.NewVersion(8, 0, 0)
	segment := index.NewSegmentInfo(dir, version, minVersion, "0", 10000,
		false, nil, map[string]string{}, []byte("1"), map[string]string{}, nil)

	fieldInfos := index.NewFieldInfos([]*document.FieldInfo{&document.FieldInfo{}})

	writeState := index.NewSegmentWriteState(dir, segment, fieldInfos, index.NewBufferedUpdates(), nil)

	writer, err := format.FieldsWriter(writeState)
	if err != nil {
		panic(err)
	}

	fieldInfo := document.NewFieldInfo(
		"field1",
		1,
		false,
		false,
		true,
		document.INDEX_OPTIONS_DOCS_AND_FREQS_AND_POSITIONS_AND_OFFSETS,
		document.DOC_VALUES_TYPE_NONE,
		-1,
		map[string]string{},
		2,
		0,
		4,
		true,
	)

	readState := index.NewSegmentReadState(dir, segment, fieldInfos, nil, "")

	reader, err := simpletext.NewPointsReader(readState)
	if err != nil {
		panic(err)
	}

	err = writer.WriteField(fieldInfo, reader)
	if err != nil {
		panic(err)
	}

}
