package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applyDepthFilter(t *testing.T) {
	r := []string{
		"",
		" ├── ",
		" ├── ",
		" │    └── ",
		" ├── ",
		" │    ├── ",
		" │    ├── ",
		" │    │    ├── ",
		" │    │    │    └── ",
	}
	c := []string{
		"parent",
		"child1",
		"child2",
		"grand-child1_child_of_child2",
		"child3",
		"grand-child2_child1_of_child3",
		"grand-child3_child2_of_child3",
		"great-grand-child1_child_of_grandchild3",
		"great-great-grand-child1_child_of_great-grand-child1",
	}

	var tests = map[string]struct {
		visualizeTrimmed bool
		wantL7Route      string
		wantL7Content    string
		wantL7View       bool
	}{
		"visualized": {
			visualizeTrimmed: true,
			wantL7Route:      " │    │    └── ",
			wantL7Content:    depthMarker,
			wantL7View:       true,
		},
		"not_visualized": {
			visualizeTrimmed: false,
			wantL7Route:      r[7],
			wantL7Content:    c[7],
			wantL7View:       false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			l0 := routeTreeLine{route: r[0], content: c[0], view: true}
			l1 := routeTreeLine{route: r[1], content: c[1], view: true, parent: &l0}
			l2 := routeTreeLine{route: r[2], content: c[2], view: true, parent: &l0}
			l3 := routeTreeLine{route: r[3], content: c[3], view: true, parent: &l2}
			l4 := routeTreeLine{route: r[4], content: c[4], view: true, parent: &l0}
			l5 := routeTreeLine{route: r[5], content: c[5], view: true, parent: &l4}
			l6 := routeTreeLine{route: r[6], content: c[6], view: true, parent: &l4}
			l7 := routeTreeLine{route: r[7], content: c[7], view: true, parent: &l6}
			l8 := routeTreeLine{route: r[8], content: c[8], view: true, parent: &l7}
			ls := &routeTreeLines{&l0, &l1, &l2, &l3, &l4, &l5, &l6, &l7, &l8}
			// act
			got := ls.applyDepthFilter(2, tc.visualizeTrimmed)
			// assert
			assert.Equal(2, got)
			assert.Equal(r[0], l0.route)
			assert.Equal(r[1], l1.route)
			assert.Equal(r[2], l2.route)
			assert.Equal(r[3], l3.route)
			assert.Equal(r[4], l4.route)
			assert.Equal(r[5], l5.route)
			assert.Equal(r[6], l6.route)
			assert.Equal(tc.wantL7Route, l7.route)
			assert.Equal(r[8], l8.route)
			assert.Equal(c[0], l0.content)
			assert.Equal(c[1], l1.content)
			assert.Equal(c[2], l2.content)
			assert.Equal(c[3], l3.content)
			assert.Equal(c[4], l4.content)
			assert.Equal(c[5], l5.content)
			assert.Equal(c[6], l6.content)
			assert.Equal(tc.wantL7Content, l7.content)
			assert.Equal(c[8], l8.content)
			assert.True(l0.view)
			assert.True(l1.view)
			assert.True(l2.view)
			assert.True(l3.view)
			assert.True(l4.view)
			assert.True(l5.view)
			assert.True(l6.view)
			assert.Equal(tc.wantL7View, l7.view)
			assert.False(l8.view)
		})
	}
}

func Test_applyUpgradableFilter(t *testing.T) {
	// arrange
	assert := assert.New(t)
	invisible := routeTreeLine{content: "invisible"}
	parentOfVisible := routeTreeLine{content: "parent_of_visible", view: true} // parents of visible children are still visible
	visible := routeTreeLine{content: "visible", view: true, parent: &parentOfVisible}
	upgradeInvisible := routeTreeLine{content: "upgrade_invisible [v1.2.3]", view: false}
	parentOfUpgradable := routeTreeLine{content: "parent_of_upgrade", view: true} // parents of visible children are still visible
	upgradeVisible := routeTreeLine{content: "upgrade_visible [v3.2.1]", view: true, parent: &parentOfUpgradable}
	dotsAfterUpgrade := routeTreeLine{content: depthMarker, view: true, parent: &upgradeVisible}
	ls := &routeTreeLines{
		&invisible,
		&parentOfVisible,
		&visible,
		&upgradeInvisible,
		&parentOfUpgradable,
		&upgradeVisible,
		&dotsAfterUpgrade,
	}
	// act
	got := ls.applyUpgradableFilter(true)
	// assert
	assert.Equal(2, got)
	assert.Equal(false, invisible.view)
	assert.Equal(false, parentOfVisible.view)
	assert.Equal(false, visible.view)
	assert.Equal(false, upgradeInvisible.view)
	assert.Equal(true, parentOfUpgradable.view)
	assert.Equal(true, upgradeVisible.view)
	assert.Equal(true, dotsAfterUpgrade.view)
}

func Test_applyColors(t *testing.T) {
	// arrange
	assert := assert.New(t)
	invisible := routeTreeLine{content: "invisible"}
	visible := routeTreeLine{content: "visible", view: true}
	upgrade := routeTreeLine{content: "upgrade [v1.2.3]", view: true}
	ls := &routeTreeLines{&invisible, &visible, &upgrade}
	// act
	ls.applyColors()
	// assert
	assert.Equal("", invisible.color)
	assert.Equal("", visible.color)
	assert.Equal(colorYellow, upgrade.color)
}
