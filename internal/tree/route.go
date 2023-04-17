package tree

import (
	"fmt"
	"strings"
)

const (
	depthMarker   = "..."
	upgradeMarker = "["
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m" // bad too recognize on dark background
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type routeTreeLine struct {
	color   string
	route   string
	content string
	view    bool
	parent  *routeTreeLine
}

type routeTreeLines []*routeTreeLine

type routeTree struct {
	lines *routeTreeLines
}

var colorDefinition = map[string]string{upgradeMarker: colorYellow}

// printRoute prints out the tree as route to the console
func (t *tree) printRoute() {
	rt := routeTree{lines: &routeTreeLines{}}

	// evaluate data and apply filters
	rt.processItem("", *t.rootItem, nil)
	filtered := rt.lines.applyDepthFilter(t.depth, t.visualizeTrimmed)
	if !t.showAll {
		filtered = filtered + rt.lines.applyUpgradableFilter(t.visualizeTrimmed)
	}
	if t.colored {
		rt.lines.applyColors()
	}
	// ensure at least root is shown
	if len(*rt.lines) > 0 {
		(*rt.lines)[0].view = true
	}

	// prepare output
	cReset := ""
	if t.colored {
		cReset = colorReset
	}
	var toPrint []string
	for _, l := range *rt.lines {
		if l.view {
			toPrint = append(toPrint, fmt.Sprintf("%s%s%s%s", l.route, l.color, l.content, cReset))
		}
	}

	fmt.Printf("dependency tree with depth %d for package: %s", t.depth, t.rootItem.module.Name())
	if filtered > 0 {
		fmt.Printf(", least %d trimmed item(s)", filtered)
	}
	var bracketText []string
	if !t.visualizeTrimmed {
		bracketText = append(bracketText, "no visualization for trimmed tree")
	}
	if !t.showAll {
		bracketText = append(bracketText, "only upgradable items with parents")
	}
	if len(bracketText) > 0 {
		fmt.Printf(" (%s)", strings.Join(bracketText, ", "))
	}
	fmt.Printf("\n\n")
	fmt.Println(strings.Join(toPrint, "\n"))
}

// processItem add lines with route sign etc. starting at the given item
func (rt *routeTree) processItem(route string, item treeItem, parentLine *routeTreeLine) int {
	if (len([]rune(string(route)))+1)/5 > maxDepth {
		return -1
	}
	line := rt.createAndAddLine(route, item, parentLine)

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
			if rt.processItem(corner, *child, line) == -1 {
				break
			}
		}
	}

	return 0
}

// createAndAddLine creates the line object with route sign etc. and add it to the lines list
func (rt *routeTree) createAndAddLine(route string, item treeItem, parentLine *routeTreeLine) *routeTreeLine {
	lineContent := item.module.Name()
	if len(item.module.GoVersion) > 0 {
		lineContent = fmt.Sprintf("%s (go%s)", lineContent, item.module.GoVersion)
	}
	updateModule := item.module.GetUpdateModule()
	if updateModule != nil {
		lineContent = fmt.Sprintf("%s => %s%s]", lineContent, upgradeMarker, updateModule.Version)
		if len(updateModule.GoVersion) > 0 {
			lineContent = fmt.Sprintf("%s (go%s)", lineContent, updateModule.GoVersion)
		}
	}
	return rt.lines.add(route, lineContent, parentLine)
}

// add creates a new instance with the given content and add it to the list
func (ls *routeTreeLines) add(route, content string, parentLine *routeTreeLine) *routeTreeLine {
	l := &routeTreeLine{route: route, content: content, view: true, parent: parentLine}
	*ls = append(*ls, l)
	return l
}

// applyDepthFilter mark lines to show on later print to console according to given depth
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

// applyUpgradableFilter set lines without upgrade information to invisible
func (ls *routeTreeLines) applyUpgradableFilter(visualizeTrimmed bool) (countFiltered int) {
	checkNextForDots := false
	for _, l := range *ls {
		// not viewables keep untouched
		if !l.view {
			continue
		}
		if visualizeTrimmed && checkNextForDots && strings.Contains(l.content, depthMarker) {
			// keep this dots active, but not the next
			checkNextForDots = false
			continue
		}

		// set all not upgradable to not viewable
		if !strings.Contains(l.content, upgradeMarker) {
			if l.view {
				l.view = false
				countFiltered++
			}
			continue
		}

		// work for upgradable item
		checkNextForDots = true

		// recursive parent activation of viewable children, independent of upgrade status
		p := l.parent
		for {
			if p == nil {
				break
			}
			if !p.view {
				p.view = true
				countFiltered--
			}
			p = p.parent
		}
	}
	return
}

// applyColors colorize lines, which contains defined marker
func (ls *routeTreeLines) applyColors() {
	for _, l := range *ls {
		if !l.view {
			continue
		}
		for marker, color := range colorDefinition {
			if strings.Contains(l.content, marker) {
				l.color = color
			}
		}
	}
}
