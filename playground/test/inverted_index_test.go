package test

import (
	"fmt"
	"search-engine/playground"
	"testing"
)

func TestBuildInvertedIndex(t *testing.T) {
	docs := []*playground.Doc{
		&playground.Doc{1, []string{"go", "数据结构"}},
		&playground.Doc{2, []string{"go", "数据库"}},
	}

	invertedIndex := playground.BuildInvertIndex(docs)

	for key, value := range invertedIndex {
		fmt.Println(key, value)
	}

}
