package tree

import (
	"fmt"
	"strings"
)

type routeTreeLine struct {
	route   string
	content string
	view    bool
}

type routeTreeLines []*routeTreeLine

type routeTree struct {
	lines *routeTreeLines
}

// Print out the tree as route to the console
func (t *tree) printRoute() {
	if t.rootItem == nil {
		return
	}

	rt := routeTree{lines: &routeTreeLines{}}

	// evaluate data and apply filters
	rt.processItem("", *t.rootItem)
	depthFiltered := rt.lines.applyDepthFilter(t.depth, t.visualizeTrimmed)

	// prepare output
	var toPrint []string
	for _, l := range *rt.lines {
		if l.view {
			toPrint = append(toPrint, fmt.Sprintf("%s%s", l.route, l.content))
		}
	}

	fmt.Printf("dependency tree with depth %d for package: %s", t.depth, t.rootItem.info.Name())
	if depthFiltered > 0 {
		fmt.Printf(", least %d trimmed item(s)", depthFiltered)
	}
	if !t.visualizeTrimmed {
		fmt.Printf(" (no visualization for trimmed tree)")
	}
	fmt.Printf("\n\n")

	fmt.Println(strings.Join(toPrint, "\n"))
}

func (rt *routeTree) processItem(route string, item treeItem) int {
	if (len([]rune(string(route)))+1)/5 > maxDepth {
		return -1
	}
	rt.createAndAddLine(route, item)

	if len(item.children) > 0 {
		childPath := route
		childPath = strings.Replace(childPath, "└", " ", -1)
		childPath = strings.Replace(childPath, "├", "│", -1)
		childPath = strings.Replace(childPath, "─", " ", -1)
		childPath = strings.Replace(childPath, "┌", "│", -1)

		childLen := len(item.children)
		for i, child := range item.children {
			corner := ""
			if i == childLen-1 {
				corner = childPath + " └── "
			} else {
				corner = childPath + " ├── "
			}
			if rt.processItem(corner, *child) == -1 {
				break
			}
		}
	}

	return 0
}

func (ls *routeTreeLines) applyDepthFilter(printDepth int, visualizeTrimmed bool) (countFiltered int) {
	skipNext := false
	for _, l := range *ls {
		if (len([]rune(string(l.route)))+3)/5 > printDepth {
			countFiltered++
			if !skipNext {
				if visualizeTrimmed {
					l.route = strings.Replace(l.route, "├", "└", -1)
					l.content = depthMarker
				} else {
					l.view = false
				}
				skipNext = true
			} else {
				l.view = false
			}
		} else {
			skipNext = false
		}
	}
	return
}

func (rt *routeTree) createAndAddLine(route string, item treeItem) {
	lineContent := item.info.Name()
	rt.lines.add(route, lineContent)
}

func (ls *routeTreeLines) add(route, content string) *routeTreeLine {
	l := &routeTreeLine{route: route, content: content, view: true}
	*ls = append(*ls, l)
	return l
}
