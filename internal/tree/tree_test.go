package tree

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vc60er/deptree/internal/moduleinfo"
)

func TestNewTree(t *testing.T) {
	var tests = map[string]struct {
		depth            int
		visualizeTrimmed bool
		showAll          bool
		colored          bool
		wantDepth        int
	}{
		"max_depth": {depth: maxDepth + 1, colored: true, wantDepth: maxDepth},
		"depth":     {depth: maxDepth - 3, visualizeTrimmed: true, wantDepth: maxDepth - 3},
		"showAll":   {depth: 0, showAll: true, wantDepth: 1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			require := require.New(t)
			i := moduleinfo.NewInfo()
			// act
			got := NewTree(tc.depth, tc.visualizeTrimmed, tc.showAll, tc.colored, *i)
			// assert
			require.NotNil(got.items)
			assert.Equal(tc.wantDepth, got.depth)
			assert.Equal(tc.visualizeTrimmed, got.visualizeTrimmed)
			assert.Equal(tc.showAll, got.showAll)
			assert.Equal(tc.colored, got.colored)
		})
	}
}

func TestFill(t *testing.T) {
	// arrange
	assert := assert.New(t)
	require := require.New(t)
	i := moduleinfo.NewInfo()
	got := tree{items: make(map[string]*treeItem), modInfo: *i}
	// act
	file := graphFile("../../test/data/graphfile_small.txt")
	defer file.Close()
	got.Fill(file)
	// assert
	require.NotNil(got.items)
	assert.Equal(13, len(got.items))
	require.NotNil(got.items["github.com/vc60er/deptree"])
	assert.Equal(1, len(got.items["github.com/vc60er/deptree"].children))
	require.NotNil(got.items["github.com/stretchr/testify@v1.8.2"])
	assert.Equal(4, len(got.items["github.com/stretchr/testify@v1.8.2"].children))
}

func graphFile(graphFile string) *os.File {
	var err error
	if graphFile, err = filepath.Abs(graphFile); err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(graphFile)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
