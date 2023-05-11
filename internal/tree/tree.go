package tree

import (
	"bufio"
	"io"
	"log"
	"strings"

	"github.com/vc60er/deptree/internal/moduleinfo"
	"github.com/vc60er/deptree/internal/verbose"
)

const maxDepth = 20

type treeItem struct {
	module   *moduleinfo.Module
	children []*treeItem
}

type tree struct {
	verbose          verbose.Verbose
	depth            int
	showDroppedChild bool
	visualizeTrimmed bool
	showAll          bool
	colored          bool
	modInfo          moduleinfo.Info
	items            map[string]*treeItem
	rootItem         *treeItem
}

// NewTree creates a new instance for tree visualization
func NewTree(verbose verbose.Verbose, depth int,
	showDroppedChild, visualizeTrimmed, showAll, colored bool, modInfo moduleinfo.Info) *tree {
	if depth > maxDepth {
		depth = maxDepth
	}
	if depth < 1 {
		depth = 1
	}
	t := tree{
		verbose:          verbose,
		depth:            depth,
		showDroppedChild: showDroppedChild,
		visualizeTrimmed: visualizeTrimmed,
		showAll:          showAll,
		colored:          colored,
		modInfo:          modInfo,
		items:            make(map[string]*treeItem),
	}
	return &t
}

func (t *tree) Fill(file io.Reader) {
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		ss := strings.Split(string(line), " ")
		const lineParts = 2
		if len(ss) != lineParts {
			log.Fatal("error input")
		}

		t.addItem(ss[0], ss[1])
	}
	t.verbose.Log1f("%d graph items collected", len(t.items))
}

// Print prints the collected information to command line (STDOUT) in the given format.
func (t *tree) Print(asJSON bool) {
	if t.rootItem == nil {
		return
	}

	if asJSON {
		t.printJSON()
		return
	}

	t.printRoute()
}

// addItem creates a new instance and add it to the list of items or enhance an existing item by children.
func (t *tree) addItem(name string, childName string) *treeItem {
	item, ok := t.items[name]
	if !ok {
		item = &treeItem{
			module: t.modInfo.GetModuleAddIfEmpty(name),
		}
		t.items[name] = item
	}

	if t.rootItem == nil {
		t.rootItem = item
	}

	if len(childName) > 0 {
		if name == childName {
			log.Fatalf("circular dependency for %s", name)
		}
		if contains(item.children, childName) {
			log.Fatalf("try to add duplicate children %s", childName)
		}
		item.children = append(item.children, t.addItem(childName, ""))
	}

	return item
}

func contains(items []*treeItem, searchName string) bool {
	for _, item := range items {
		if item.module.Name() == searchName {
			return true
		}
	}
	return false
}
