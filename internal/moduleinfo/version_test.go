package moduleinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_higherOrEqualVersion(t *testing.T) {
	var tests = map[string]struct {
		vam *Module
		vbm *Module
	}{
		"empty_equal": {vam: &Module{}, vbm: &Module{}},
		"semantic_equal": {
			vam: &Module{Version: "v2.1.0"},
			vbm: &Module{Version: "v2.1.0"},
		},
		"commit_equal": {
			vam: &Module{Version: "v0.0.0-20190916202348-b4ddaad3f8a3"},
			vbm: &Module{Version: "v0.0.0-20190916202348-b4ddaad3f8a3"},
		},
		"semantic_major": {
			vam: &Module{Version: "v3.0.0"},
			vbm: &Module{Version: "v2.0.0"},
		},
		"semantic_minor": {
			vam: &Module{Version: "v2.1.1"},
			vbm: &Module{Version: "v2.0.1"},
		},
		"semantic_patch": {
			vam: &Module{Version: "v2.1.1"},
			vbm: &Module{Version: "v2.1.0"},
		},
		"commit": {
			vam: &Module{Version: "v0.0.0-20220715151400-c0bba94af5f8"},
			vbm: &Module{Version: "v0.0.0-20190916202348-b4ddaad3f8a3"},
		},
		"one_is_empty": {vam: &Module{Version: "v0.1.0"}, vbm: &Module{}},
		"one_as_commit": {
			vam: &Module{Version: "v0.0.0-20220715151400-c0bba94af5f8"},
			vbm: &Module{Version: "v0.0.0"},
		},
		"one_as_rc": {
			vam: &Module{Version: "v4.3.0+incompatible"},
			vbm: &Module{Version: "v4.3.0"},
		},
		"one_as_incompatible": {
			vam: &Module{Version: "v1.26.0-rc.1"},
			vbm: &Module{Version: "v1.26.0"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			// act
			got1 := higherOrEqualVersion(tc.vam, tc.vbm)
			got2 := higherOrEqualVersion(tc.vbm, tc.vam)
			// assert
			assert.Equal(tc.vam, got1)
			assert.Equal(tc.vam, got2)
		})
	}
}
