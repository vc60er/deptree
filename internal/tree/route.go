package tree

import (
	"fmt"
	"strings"

	"github.com/vc60er/deptree/internal/verbose"
)

const (
	depthMarker            = "..."
	upgradeMarker          = "["
	alreadyProcessedMarker = "see <"
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

type routeTreeLines []*routeTreeLine

type routeTreeLine struct {
	color             string
	route             string
	level             int
	link              string
	name              string
	additionalContent string
	view              bool
	parent            *routeTreeLine
	children          routeTreeLines
}

type routeTree struct {
	verbose          verbose.Verbose
	maxDepth         int
	showDroppedChild bool
	lines            *routeTreeLines
	mostChildren     treeItem
}

var colorDefinition = map[string]string{upgradeMarker: colorYellow, alreadyProcessedMarker: colorCyan}

// printRoute prints out the tree as route to the console
func (t *tree) printRoute() {
	rt := routeTree{
		verbose:          t.verbose,
		maxDepth:         t.depth,
		showDroppedChild: t.showDroppedChild,
		lines:            &routeTreeLines{},
	}
	if t.visualizeTrimmed {
		// to prepare the route for "..." lines, we need one level more
		rt.maxDepth++
	}

	// evaluate data and apply filters
	rt.processItem(*t.rootItem, nil)
	t.verbose.Log1f("%d tree items processed", len(*rt.lines))
	t.verbose.Log1f("item with most children %s (%d)", rt.mostChildren.module.Name(), len(rt.mostChildren.children))

	filtered := rt.lines.applyDepthFilter(t.depth, t.visualizeTrimmed)
	if filtered > 0 {
		rt.lines = rt.lines.cleanInvisible()
	}

	if !t.showAll {
		f := rt.lines.applyUpgradableFilter(t.visualizeTrimmed)
		if f > 0 {
			rt.lines = rt.lines.cleanInvisible()
			filtered += f
		}
	}

	var duplicateChildTrees int
	if !t.showDroppedChild {
		f, l := rt.lines.applyAlreadyProcessedFilter(t.showAll)
		duplicateChildTrees = l
		if f > 0 {
			rt.lines = rt.lines.cleanInvisible()
			filtered += f
		}
	}

	if t.colored {
		rt.lines.applyColors()
	}

	// prepare output (all invisible lines are already removed from list)
	if len(*rt.lines) > 0 {
		firstLine := (*rt.lines)[0]
		firstLine.view = true // ensure at least root is shown
		createRoutes(firstLine, "")
	}

	var cReset string
	if t.colored {
		cReset = colorReset
	}

	var toPrint []string
	for _, l := range *rt.lines {
		toPrint = append(toPrint, fmt.Sprintf("%s%s%s%s%s%s", l.route, l.link, l.color, l.name, l.additionalContent, cReset))
	}

	txt := fmt.Sprintf("dependency tree with depth %d for package: %s", t.depth, t.rootItem.module.Name())
	if filtered > 0 {
		txt = fmt.Sprintf("%s, least %d trimmed item(s)", txt, filtered)
	}
	fmt.Println(txt)

	if !t.visualizeTrimmed {
		fmt.Println("* no visualization for trimmed tree (-t not set)")
	}
	if !t.showAll {
		fmt.Println("* only upgradable items with parents are shown (-a not set)")
	}
	if !t.showDroppedChild {
		if !t.showAll {
			fmt.Println("* duplicate children not shown (-a not set)")
		} else {
			fmt.Println("* tree of duplicate children not drawn (-f not set)")
		}
	}
	if duplicateChildTrees > 0 {
		fmt.Printf("* %d duplicate child trees replaced by links\n", duplicateChildTrees)
	}
	fmt.Println("")
	fmt.Println(strings.Join(toPrint, "\n"))
}

// processItem add lines with route sign etc. starting at the given item
func (rt *routeTree) processItem(item treeItem, parentLine *routeTreeLine) {
	var branchName string
	var level int
	if parentLine != nil {
		branchName = parentLine.name
		level = parentLine.level + 1
	}
	if level > rt.maxDepth {
		rt.verbose.Log3f("maxDepth reached for %s", item.module.Name())
		return
	}

	name := item.module.Name()
	childLen := len(item.children)
	rt.verbose.Log2f("process branch '%s', level %d, item %s, children: %d",
		branchName, level, name, childLen)

	line := rt.createAndAddLine(level, item, parentLine)

	if childLen > 0 {
		if len(rt.mostChildren.children) < childLen {
			rt.mostChildren = item
		}
		// we only collect children on new branch, if found at a higher position (lower level)
		if !rt.showDroppedChild {
			ho := rt.lines.highestOccurrence(name)
			if ho.branch() != branchName && ho.level <= level {
				rt.verbose.Log2f("%d children dropped at branch '%s' level %d for %s",
					childLen, branchName, level, item.module.Name())
				return
			}
		}
		for _, child := range item.children {
			rt.processItem(*child, line)
		}
	}
}

// createAndAddLine creates the line object with route sign etc. and add it to the lines list
func (rt *routeTree) createAndAddLine(level int, item treeItem, parentLine *routeTreeLine) *routeTreeLine {
	var addContent string

	if len(item.module.GoVersion) > 0 {
		addContent = fmt.Sprintf(" (go%s)", item.module.GoVersion)
	}

	updateModule := item.module.GetUpdateModule()
	if updateModule != nil {
		addContent = fmt.Sprintf("%s => %s%s]", addContent, upgradeMarker, updateModule.Version)
		if len(updateModule.GoVersion) > 0 {
			addContent = fmt.Sprintf("%s (go%s)", addContent, updateModule.GoVersion)
		}
	}

	return rt.lines.add(level, item.module.Name(), addContent, parentLine)
}

func (ls *routeTreeLines) highestOccurrence(name string) *routeTreeLine {
	var highestOccurence *routeTreeLine
	for _, l := range *ls {
		if l.name == name && (highestOccurence == nil || highestOccurence.level > l.level) {
			highestOccurence = l
		}
	}
	return highestOccurence
}

// add creates a new instance with the given content and add it to the list
func (ls *routeTreeLines) add(level int, name, additionalContent string, parentLine *routeTreeLine) *routeTreeLine {
	l := routeTreeLine{
		level:             level,
		name:              name,
		additionalContent: additionalContent,
		view:              true,
		parent:            parentLine,
	}
	*ls = append(*ls, &l)
	if parentLine != nil {
		parentLine.children = append(parentLine.children, &l)
	}
	return &l
}

// applyDepthFilter mark lines to show on later print to console according to given depth
func (ls *routeTreeLines) applyDepthFilter(printDepth int, visualizeTrimmed bool) int {
	var skipNext bool
	var countFiltered int
	for _, l := range *ls {
		if l.level > printDepth {
			countFiltered++
			if !skipNext {
				if visualizeTrimmed {
					l.name = ""
					if l.parent != nil {
						l.name = fmt.Sprintf("%d more ", len(l.parent.children))
					}
					l.additionalContent = depthMarker
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
	return countFiltered
}

// applyUpgradableFilter set lines without upgrade information to invisible
func (ls *routeTreeLines) applyUpgradableFilter(visualizeTrimmed bool) int {
	var checkNextForDots bool
	var countFiltered int
	for _, l := range *ls {
		// not viewables keep untouched
		if !l.view {
			continue
		}
		if visualizeTrimmed && checkNextForDots && strings.Contains(l.additionalContent, depthMarker) {
			// keep this dots active, but not the next
			checkNextForDots = false
			continue
		}

		// set all not upgradable to not viewable
		if !strings.Contains(l.additionalContent, upgradeMarker) {
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
	return countFiltered
}

// applyAlreadyProcessedFilter set lines which points to already processed children to invisible, except showAll is set,
// for showAll=true already processed children will be linked to the first occurrence with lowest level, which means:
// * this child get a link and its tree will be print completely
// * further occurrences of child tree will be linked to the first (with note the parent branch, to find quickly)
//
//nolint:nonamedreturns // OK here for better differentiation of both int returns
func (ls *routeTreeLines) applyAlreadyProcessedFilter(showAll bool) (countFiltered int, createdLinks int) {
	// if not showAll, set all occurrences to invisible, except the highest and return
	if !showAll {
		for _, l := range *ls {
			// not viewables keep untouched and at root level all items occur the first time
			if !l.view || l.parent == nil {
				continue
			}

			all := ls.allOccurrences(l.name)
			if len(all) <= 1 {
				// there is only one occurrence (the first), so keep it as is
				continue
			}

			highestOccurrence := all.highestOccurrence(l.name)
			if highestOccurrence.branch() != l.parent.name {
				// this is not the highest occurrence, set children to invisible
				for _, c := range l.children {
					if c.view {
						c.view = false
						countFiltered++
					}
				}
				continue
			}
		}
		return countFiltered, 0
	}

	// if all occurrences should be shown (showAll==true) and if there are visible children, we need some links
	linkNumbers := make(map[*routeTreeLine]int)
	for _, l := range *ls {
		// not viewables keep untouched and at root level all items occur the first time
		if !l.view || l.parent == nil {
			continue
		}

		// same items on same or different branch can lead to complex/indirect circular dependencies, so we don't draw it
		// each time, but replace it with a link
		all := ls.allOccurrences(l.name)
		if len(all) <= 1 {
			// there is only one occurrence (the first), so keep it as is
			continue
		}

		highestOccurrence := all.highestOccurrence(l.name)

		countVisibleChildren := highestOccurrence.children.countVisible()
		if countVisibleChildren <= 0 {
			// there is nothing to do if no visible children exist at first occurrence
			continue
		}

		firstOccurrenceBranch := highestOccurrence.branch()
		if firstOccurrenceBranch != l.parent.name {
			// this is not the first occurrence, so we need a link for the line content
			linkNumber, ok := linkNumbers[highestOccurrence]
			if !ok {
				// this is the first link request, so we create a new link
				linkNumber = len(linkNumbers) + 1
				linkNumbers[highestOccurrence] = linkNumber
				highestOccurrence.link = fmt.Sprintf("<%d> ", linkNumber)
			}
			l.additionalContent = fmt.Sprintf("%s, see <%d> for %d children (branch '%s' level %d)",
				l.additionalContent, linkNumber, countVisibleChildren, firstOccurrenceBranch, highestOccurrence.level)
			// afterwards all children will be switched off
			for _, c := range l.children {
				if c.view {
					c.view = false
					countFiltered++
				}
			}
		}
	}
	return countFiltered, len(linkNumbers)
}

func (ls *routeTreeLines) allOccurrences(name string) routeTreeLines {
	var occurrences []*routeTreeLine
	for _, l := range *ls {
		if l.name == name {
			occurrences = append(occurrences, l)
		}
	}
	return occurrences
}

func (ls *routeTreeLines) countVisible() int {
	var count int
	for _, l := range *ls {
		if l.view {
			count++
		}
	}
	return count
}

func (ls *routeTreeLines) cleanInvisible() *routeTreeLines {
	var newLines routeTreeLines
	for _, l := range *ls {
		if l.view {
			// do not clean children recursive
			var newChildren routeTreeLines
			for _, c := range l.children {
				if c.view {
					newChildren = append(newChildren, c)
				}
			}
			l.children = newChildren
			newLines = append(newLines, l)
		}
	}
	return &newLines
}

// applyColors colorize lines, which contains defined marker
func (ls *routeTreeLines) applyColors() {
	for _, l := range *ls {
		if !l.view {
			continue
		}
		if l.link != "" {
			l.link = fmt.Sprintf("%s%s%s", colorDefinition[alreadyProcessedMarker], l.link, colorReset)
		}
		for marker, color := range colorDefinition {
			if strings.Contains(l.additionalContent, marker) {
				l.color = color
				if marker == alreadyProcessedMarker {
					// has precedence over all other colors
					break
				}
			}
		}
	}
}

func (l *routeTreeLine) branch() string {
	if l.parent == nil {
		return ""
	}
	return l.parent.name
}

func createRoutes(line *routeTreeLine, route string) {
	line.route = route
	childLen := len(line.children)
	if childLen > 0 {
		childPath := route
		childPath = strings.ReplaceAll(childPath, "└", " ")
		childPath = strings.ReplaceAll(childPath, "├", "│")
		childPath = strings.ReplaceAll(childPath, "─", " ")
		for i, child := range line.children {
			var corner string
			if i == childLen-1 {
				corner = childPath + " └── "
			} else {
				corner = childPath + " ├── "
			}
			createRoutes(child, corner)
		}
	}
}
