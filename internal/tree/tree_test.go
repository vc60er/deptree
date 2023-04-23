package tree

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vc60er/deptree/internal/moduleinfo"
	"github.com/vc60er/deptree/internal/verbose"
)

func TestNewTree(t *testing.T) {
	var tests = map[string]struct {
		depth            int
		visualizeTrimmed bool
		showDroppedChild bool
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
			v := verbose.Verbose{}
			i := moduleinfo.NewInfo(v)
			// act
			got := NewTree(v, tc.depth, tc.showDroppedChild, tc.visualizeTrimmed, tc.showAll, tc.colored, *i)
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
	v := verbose.Verbose{}
	i := moduleinfo.NewInfo(v)
	got := tree{items: make(map[string]*treeItem), modInfo: *i}
	// act
	file := graphFile("../../test/data/graphfile_small.txt")
	defer file.Close()
	got.Fill(file)
	// assert
	require.NotNil(got.items)
	assert.Equal(13, len(got.items))
	// assert root level 0
	root := got.items["github.com/vc60er/deptree"]
	require.NotNil(root)
	assert.Equal(1, len(root.children))
	// assert level 1
	level1_1 := got.items["github.com/stretchr/testify@v1.8.2"]
	require.NotNil(level1_1)
	assert.Equal(4, len(level1_1.children))
	// assert level 2
	level2_1 := got.items["github.com/davecgh/go-spew@v1.1.1"]
	require.NotNil(level2_1)
	assert.Equal(0, len(level2_1.children))
	level2_2 := got.items["github.com/pmezard/go-difflib@v1.0.0"]
	require.NotNil(level2_2)
	assert.Equal(0, len(level2_2.children))
	level2_3 := got.items["github.com/stretchr/objx@v0.5.0"]
	require.NotNil(level2_3)
	assert.Equal(1, len(level2_3.children))
	level2_4 := got.items["gopkg.in/yaml.v3@v3.0.1"]
	require.NotNil(level2_4)
	assert.Equal(1, len(level2_4.children))
	// assert level 3
	level3_1 := got.items["github.com/stretchr/testify@v1.8.0"]
	require.NotNil(level3_1)
	assert.Equal(4, len(level3_1.children))
	level3_2 := got.items["gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405"]
	require.NotNil(level3_2)
	assert.Equal(0, len(level3_2.children))
	// assert level 4
	level4_1 := got.items["github.com/stretchr/objx@v0.4.0"]
	require.NotNil(level4_1)
	assert.Equal(2, len(level4_1.children))
	// assert level 5
	level5_1 := got.items["github.com/stretchr/testify@v1.7.1"]
	require.NotNil(level5_1)
	assert.Equal(4, len(level5_1.children))
	// assert level 6
	level6_1 := got.items["github.com/davecgh/go-spew@v1.1.0"]
	require.NotNil(level6_1)
	assert.Equal(0, len(level6_1.children))
	level6_2 := got.items["github.com/stretchr/objx@v0.1.0"]
	require.NotNil(level6_2)
	assert.Equal(0, len(level6_2.children))
	level6_3 := got.items["gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c"]
	require.NotNil(level6_3)
	assert.Equal(1, len(level6_3.children))
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
