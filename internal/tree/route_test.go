package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add(t *testing.T) {
	// arrange
	assert := assert.New(t)
	invisible := routeTreeLine{name: "invisible"}
	visible := routeTreeLine{name: "visible", view: true}
	ls := &routeTreeLines{&invisible, &visible}
	// act
	ls.add(5, "to add", "extra content", &visible)
	// assert
	assert.Equal(3, len(*ls))
	newLine := (*ls)[2]
	assert.Equal("", newLine.color)
	assert.Equal("", newLine.route)
	assert.Equal(5, newLine.level)
	assert.Equal("", newLine.link)
	assert.Equal("to add", newLine.name)
	assert.Equal("extra content", newLine.additionalContent)
	assert.True(newLine.view)
	assert.Equal(visible, *newLine.parent)
	assert.Equal(newLine, visible.children[0])
}

func Test_applyDepthFilter(t *testing.T) {
	n := []string{
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
		wantL7Name       string
		wantL7AddContent string
		wantL7View       bool
	}{
		"visualized": {
			visualizeTrimmed: true,
			wantL7Name:       "1 more ",
			wantL7AddContent: depthMarker,
			wantL7View:       true,
		},
		"not_visualized": {
			visualizeTrimmed: false,
			wantL7Name:       n[7],
			wantL7AddContent: "",
			wantL7View:       false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			l0 := routeTreeLine{name: n[0], view: true}
			l1 := routeTreeLine{name: n[1], view: true, parent: &l0}
			l2 := routeTreeLine{name: n[2], view: true, parent: &l0}
			l3 := routeTreeLine{name: n[3], view: true, parent: &l2}
			l4 := routeTreeLine{name: n[4], view: true, parent: &l0}
			l5 := routeTreeLine{name: n[5], view: true, parent: &l4}
			l6 := routeTreeLine{name: n[6], view: true, parent: &l4}
			l7 := routeTreeLine{name: n[7], view: true, parent: &l6}
			l8 := routeTreeLine{name: n[8], view: true, parent: &l7}
			l0.children = routeTreeLines{&l1, &l2, &l4}
			l2.children = routeTreeLines{&l3}
			l4.children = routeTreeLines{&l5, &l6}
			l6.children = routeTreeLines{&l7}
			l7.children = routeTreeLines{&l8}
			ls := &routeTreeLines{&l0, &l1, &l2, &l3, &l4, &l5, &l6, &l7, &l8}
			setLevels(ls)
			// act
			got := ls.applyDepthFilter(2, tc.visualizeTrimmed)
			// assert
			assert.Equal(2, got)
			assert.Equal(n[0], l0.name)
			assert.Equal(n[1], l1.name)
			assert.Equal(n[2], l2.name)
			assert.Equal(n[3], l3.name)
			assert.Equal(n[4], l4.name)
			assert.Equal(n[5], l5.name)
			assert.Equal(n[6], l6.name)
			assert.Equal(tc.wantL7Name, l7.name)
			assert.Equal(tc.wantL7AddContent, l7.additionalContent)
			assert.Equal(n[8], l8.name)
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
	invisible := routeTreeLine{name: "invisible"}
	parentOfVisible := routeTreeLine{name: "parent_of_visible", view: true} // parents of visible children are still visible
	visible := routeTreeLine{name: "visible", view: true, parent: &parentOfVisible}
	upgradeInvisible := routeTreeLine{name: "upgrade_invisible", additionalContent: " [v1.2.3]", view: false}
	parentOfUpgradable := routeTreeLine{name: "parent_of_upgrade", view: true} // parents of visible children are still visible
	upgradeVisible := routeTreeLine{name: "upgrade_visible", additionalContent: " [v3.2.1]", view: true, parent: &parentOfUpgradable}
	dotsAfterUpgrade := routeTreeLine{name: "", additionalContent: depthMarker, view: true, parent: &upgradeVisible}
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

func Test_applyAlreadyProcessedFilter(t *testing.T) {
	n := []string{
		"parent_visible",
		"child1_invisible",
		"child2_visible_one_occurrence_child3_child4_child5", // no link because only one occurrence
		"child3_visible_two_occurrences_no_link",             // no link because no children
		"child4_visible_two_occurrences_no_link",             // no link because invisible children
		"child41_invisible",
		"child5_visible_two_occurrences_with_link", // for link with visible children
		"child51_visible",
	}

	var tests = map[string]struct {
		showAll          bool
		wantFiltered     int
		wantLinks        int
		wantShowL8L9L10  bool
		wantLinkL11      string
		wantAddContentL6 string
	}{
		"showAll": {
			showAll:          true,
			wantFiltered:     1,
			wantLinks:        1,
			wantShowL8L9L10:  true,
			wantLinkL11:      "<1> ",
			wantAddContentL6: ", see <1> for 1 children (branch 'parent_visible' level 1)",
		},
		"not_showAll": {
			showAll:          false,
			wantFiltered:     1,
			wantLinks:        0,
			wantShowL8L9L10:  false,
			wantLinkL11:      "",
			wantAddContentL6: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			l0 := routeTreeLine{name: n[0], view: true}
			l1 := routeTreeLine{name: n[1], view: false, parent: &l0}
			l2 := routeTreeLine{name: n[2], view: true, parent: &l0}
			l3 := routeTreeLine{name: n[3], view: true, parent: &l2}
			l4 := routeTreeLine{name: n[4], view: true, parent: &l2}
			l5 := routeTreeLine{name: n[5], view: false, parent: &l4}
			l6 := routeTreeLine{name: n[6], view: true, parent: &l2}
			l7 := routeTreeLine{name: n[7], view: true, parent: &l6}
			l8 := routeTreeLine{name: n[3], view: true, parent: &l0}
			l9 := routeTreeLine{name: n[4], view: true, parent: &l0}
			l10 := routeTreeLine{name: n[5], view: false, parent: &l9}
			l11 := routeTreeLine{name: n[6], view: true, parent: &l0}
			l12 := routeTreeLine{name: n[7], view: true, parent: &l11}
			l0.children = routeTreeLines{&l1, &l2, &l8, &l9, &l11}
			l2.children = routeTreeLines{&l3, &l4, &l6}
			l4.children = routeTreeLines{&l5}
			l6.children = routeTreeLines{&l7}
			l9.children = routeTreeLines{&l10}
			l11.children = routeTreeLines{&l12}
			ls := &routeTreeLines{&l0, &l1, &l2, &l3, &l4, &l5, &l6, &l7, &l8, &l9, &l10, &l11, &l12}
			setLevels(ls)
			// act
			gotFiltered, gotLinks := ls.applyAlreadyProcessedFilter(tc.showAll)
			// assert
			assert.Equal(tc.wantFiltered, gotFiltered)
			assert.Equal(tc.wantLinks, gotLinks)
			assert.True(l0.view)
			assert.False(l1.view)
			assert.True(l2.view)
			assert.True(l3.view)
			assert.True(l4.view)
			assert.False(l5.view)
			assert.True(l6.view)
			assert.False(l7.view, fmt.Sprintf("L7 contains '%s_%s'", l7.name, l7.additionalContent))
			assert.True(l8.view)
			assert.True(l9.view)
			assert.False(l10.view)
			assert.True(l11.view)
			assert.True(l12.view)
			assert.Equal("", l0.link)
			assert.Equal("", l1.link)
			assert.Equal("", l2.link)
			assert.Equal("", l3.link)
			assert.Equal("", l4.link)
			assert.Equal("", l5.link)
			assert.Equal("", l6.link)
			assert.Equal("", l7.link)
			assert.Equal("", l8.link)
			assert.Equal("", l9.link)
			assert.Equal("", l10.link)
			assert.Equal(tc.wantLinkL11, l11.link)
			assert.Equal("", l12.link)
			assert.Equal("", l0.additionalContent)
			assert.Equal("", l1.additionalContent)
			assert.Equal("", l2.additionalContent)
			assert.Equal("", l3.additionalContent)
			assert.Equal("", l4.additionalContent)
			assert.Equal("", l5.additionalContent)
			assert.Equal(tc.wantAddContentL6, l6.additionalContent)
			assert.Equal("", l7.additionalContent)
			assert.Equal("", l8.additionalContent)
			assert.Equal("", l9.additionalContent)
			assert.Equal("", l10.additionalContent)
			assert.Equal("", l11.additionalContent)
			assert.Equal("", l12.additionalContent)
		})
	}
}

func Test_applyColors(t *testing.T) {
	// arrange
	assert := assert.New(t)
	invisible := routeTreeLine{name: "invisible"}
	visible := routeTreeLine{name: "visible", view: true}
	upgrade := routeTreeLine{name: "upgrade", additionalContent: " [v1.2.3]", view: true}
	linked := routeTreeLine{name: "linked", additionalContent: "see <3>", view: true}
	link := routeTreeLine{link: "<3>", name: "the link", view: true}
	ls := &routeTreeLines{&invisible, &visible, &upgrade, &linked, &link}
	// act
	ls.applyColors()
	// assert
	assert.Equal("", invisible.color)
	assert.Equal("", visible.color)
	assert.Equal(colorYellow, upgrade.color)
	assert.Equal(colorCyan, linked.color)
	assert.Equal("", link.color)
	assert.Contains(link.link, colorCyan)
}

func Test_cleanInvisible(t *testing.T) {
	// arrange
	assert := assert.New(t)
	l1 := routeTreeLine{name: "to clean"}
	l2 := routeTreeLine{name: "stay", view: true}
	l3 := routeTreeLine{name: "child1 to stay", view: true}
	l4 := routeTreeLine{name: "child2 to clean"}
	l5 := routeTreeLine{name: "child3 to stay", view: true}
	l2.children = routeTreeLines{&l3, &l4, &l5}
	ls := routeTreeLines{&l1, &l2, &l3, &l4, &l5}
	// act
	got := ls.cleanInvisible()
	// assert
	assert.NotSame(ls, got)
	assert.Equal(3, len(*got))
	assert.Equal(&l2, (*got)[0])
	assert.Equal(&l3, (*got)[1])
	assert.Equal(&l5, (*got)[2])
	parent := (*got)[0]
	assert.Equal(2, len(parent.children))
	assert.Equal(&l3, parent.children[0])
	assert.Equal(&l5, parent.children[1])
}

func Test_createRoutes(t *testing.T) {
	var tests = map[string]struct {
		line     routeTreeLine
		previous string
		want     []string
	}{
		"root_no_child": {
			line: routeTreeLine{view: true},
			want: []string{""},
		},
		"level1_single_no_child": {
			line:     routeTreeLine{view: true},
			previous: " └── ",
			want: []string{
				" └── ",
			},
		},
		"level1_more_no_child": {
			line:     routeTreeLine{view: true},
			previous: " ├── ",
			want: []string{
				" ├── ",
			},
		},
		"root_one_child": {
			line: routeTreeLine{
				view: true,
				children: routeTreeLines{
					&routeTreeLine{view: true},
				},
			},
			want: []string{
				"",
				" └── ",
			},
		},
		"level1_single_one_child": {
			line: routeTreeLine{
				view: true,
				children: routeTreeLines{
					&routeTreeLine{view: true},
				},
			},
			previous: " └── ",
			want: []string{
				" └── ",
				"      └── ",
			},
		},
		"level1_more_one_child": {
			line: routeTreeLine{
				view: true,
				children: routeTreeLines{
					&routeTreeLine{view: true},
				},
			},
			previous: " ├── ",
			want: []string{
				" ├── ",
				" │    └── ",
			},
		},
		"root_two_children": {
			line: routeTreeLine{
				view: true,
				children: routeTreeLines{
					&routeTreeLine{view: true},
					&routeTreeLine{view: true},
				},
			},
			want: []string{
				"",
				" ├── ",
				" └── ",
			},
		},
		"level1_single_two_children": {
			line: routeTreeLine{
				view: true,
				children: routeTreeLines{
					&routeTreeLine{view: true},
					&routeTreeLine{view: true},
				},
			},
			previous: " └── ",
			want: []string{
				" └── ",
				"      ├── ",
				"      └── ",
			},
		},
		"level1_more_two_children": {
			line: routeTreeLine{
				view: true,
				children: routeTreeLines{
					&routeTreeLine{view: true},
					&routeTreeLine{view: true},
				},
			},
			previous: " ├── ",
			want: []string{
				" ├── ",
				" │    ├── ",
				" │    └── ",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			assert := assert.New(t)
			ls := routeTreeLines{&tc.line}
			ls = append(ls, tc.line.children...)
			// act
			createRoutes(&tc.line, tc.previous)
			// assert
			var got []string
			for _, l := range ls {
				got = append(got, l.route)
			}
			assert.Equal(tc.want, got)
		})
	}
}

func setLevels(ls *routeTreeLines) {
	for _, l := range *ls {
		setLevel(l)
	}
}

func setLevel(l *routeTreeLine) {
	var level int
	p := l.parent
	for {
		if p == nil {
			break
		}
		p = p.parent
		level++
	}

	l.level = level
}
