package moduleinfo

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInfo(t *testing.T) {
	// arrange
	assert := assert.New(t)
	// act
	got := NewInfo()
	// assert
	assert.NotNil(got.modules)
}

func TestFill(t *testing.T) {
	// arrange
	assert := assert.New(t)
	require := require.New(t)
	got := Info{modules: make(map[string]*Module)}
	// act
	got.Fill(getContent("../../test/data/upgradefile_small.txt"))
	// assert
	require.NotNil(got.modules)
	assert.Equal(7, len(got.modules))
}

func TestGetModuleAddIfEmpty(t *testing.T) {
	testModule := &Module{Path: "heiner", Version: "v1.0.0"}
	var tests = map[string]struct {
		modules   map[string]*Module
		search    string
		want      *Module
		wantCount int
	}{
		"exist": {
			modules:   map[string]*Module{"heiner@v1.0.0": testModule},
			search:    "heiner@v1.0.0",
			want:      testModule,
			wantCount: 1,
		},
		"add_new": {
			modules:   map[string]*Module{"heiner@v1.0.0": testModule},
			search:    "theo@v1.1.0",
			want:      &Module{Path: "theo", Version: "v1.1.0"},
			wantCount: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			require := require.New(t)
			i := Info{modules: tc.modules}
			// act
			got := i.GetModuleAddIfEmpty(tc.search)
			// assert
			require.NotNil(i.modules)
			assert.Equal(tc.wantCount, len(i.modules))
			assert.Equal(tc.want, got)
		})
	}
}

func TestAdjust(t *testing.T) {
	// arrange
	assert := assert.New(t)
	require := require.New(t)
	testModule1 := &Module{Path: "heiner", Version: "v1.0.0"}
	testModule2 := &Module{Path: "heiner", Version: "v1.1.0"}
	testModule3 := &Module{Path: "theo", Version: "v1.0.0"}
	got := Info{
		modules: map[string]*Module{
			"heiner@v1.0.0": testModule1,
			"heiner@v1.1.0": testModule2,
			"theo@v1.0.0":   testModule3,
		},
	}
	// act
	got.Adjust()
	// assert
	require.NotNil(got.modules)
	require.Equal(1, len(testModule1.related))
	assert.Equal(testModule2, testModule1.related[0])
	require.Equal(1, len(testModule2.related))
	assert.Equal(testModule1, testModule2.related[0])
	assert.Equal(0, len(testModule3.related))
}

func getContent(upgradeFile string) (goListCallJSONContent []byte) {
	var err error
	if upgradeFile, err = filepath.Abs(upgradeFile); err != nil {
		log.Fatal(err)
	}
	if goListCallJSONContent, err = ioutil.ReadFile(upgradeFile); err != nil {
		log.Fatal(err)
	}
	return
}
