package moduleinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUpdateModule(t *testing.T) {
	moduleV1 := &Module{Version: "1.0.0"}
	moduleV2 := &Module{Version: "2.0.0"}
	moduleV2a := &Module{Version: "1.9.9", Update: moduleV2}
	var tests = map[string]struct {
		input *Module
		want  *Module
	}{
		"no_version_no_other_module": {
			input: &Module{},
			want:  nil,
		},
		"no_version_update_module": {
			input: &Module{Update: moduleV1},
			want:  moduleV1,
		},
		"no_version_related_module": {
			input: &Module{Update: moduleV1, related: []*Module{moduleV2}},
			want:  moduleV2,
		},
		"no_version_update_module_with_related": {
			input: &Module{Update: moduleV2, related: []*Module{moduleV1}},
			want:  moduleV2,
		},
		"no_other_module": {
			input: &Module{Version: "v0.0.1"},
			want:  nil,
		},
		"update_module": {
			input: &Module{Version: "v0.0.2", Update: moduleV1},
			want:  moduleV1,
		},
		"related_module": {
			input: &Module{Version: "v0.0.3", Update: moduleV1, related: []*Module{moduleV2}},
			want:  moduleV2,
		},
		"update_module_with_related": {
			input: &Module{Version: "v0.0.4", Update: moduleV2, related: []*Module{moduleV1}},
			want:  moduleV2,
		},
		"related_module_with_update": {
			input: &Module{Version: "v0.0.5", Update: moduleV1, related: []*Module{moduleV2a}},
			want:  moduleV2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			// act
			got := tc.input.GetUpdateModule()
			// assert
			assert.Equal(tc.want, got)
		})
	}
}

func TestName(t *testing.T) {
	// arrange
	assert := assert.New(t)
	m1 := &Module{Path: "egon"}
	m2 := &Module{Path: "karl", Version: "v1.2.3"}
	// act
	got1 := m1.Name()
	got2 := m2.Name()
	// assert
	assert.Equal("egon", got1)
	assert.Equal("karl@v1.2.3", got2)
}
