package main

import (
	"encoding/binary"
	"fmt"
	"github.com/geange/lucene-go/codecs/simpletext"
	"github.com/geange/lucene-go/core/document"
	"github.com/geange/lucene-go/core/index"
	"github.com/geange/lucene-go/core/store"
	"github.com/geange/lucene-go/core/types"
	"github.com/geange/lucene-go/core/util"
)

func main() {
	//index.NewSegmentReadState()

	dir, err := store.NewNIOFSDirectory("data")
	if err != nil {
		panic(err)
	}

	format := simpletext.NewPointsFormat()

	version := util.NewVersion(8, 11, 0)
	minVersion := util.NewVersion(8, 0, 0)
	segment := index.NewSegmentInfo(dir, version, minVersion, "0", 10000,
		false, nil, map[string]string{}, []byte("1"), map[string]string{}, nil)

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
		2,
		4,
		true,
	)

	fieldInfos := index.NewFieldInfos([]*document.FieldInfo{fieldInfo})

	readState := index.NewSegmentReadState(dir, segment, fieldInfos, nil, "")

	reader, err := format.FieldsReader(readState)
	if err != nil {
		panic(err)
	}

	values, err := reader.GetValues("field1")
	if err != nil {
		panic(err)
	}

	values.Intersect(&types.BytesVisitor{
		VisitFn: func(docID int) error {
			return nil
		},
		VisitLeafFn: func(docID int, packedValue []byte) error {
			fmt.Printf("docID: %d %d,%d\n",
				docID, binary.BigEndian.Uint32(packedValue[:4]), binary.BigEndian.Uint32(packedValue[4:]))
			return nil
		},
		CompareFn: func(minPackedValue, maxPackedValue []byte) types.Relation {
			return types.CELL_CROSSES_QUERY
		},
		GrowFn: func(count int) {

		},
	})
}
