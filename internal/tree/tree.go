package tree

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/vc60er/deptree/internal/moduleinfo"
)

const maxDepth = 20

type treeItem struct {
	module   *moduleinfo.Module
	children []*treeItem
}

type tree struct {
	depth            int
	visualizeTrimmed bool
	showAll          bool
	colored          bool
	modInfo          moduleinfo.Info
	items            map[string]*treeItem
	rootItem         *treeItem
}

// NewTree creates a new instance for tree visualization
func NewTree(depth int, visualizeTrimmed, showAll, colored bool, modInfo moduleinfo.Info) *tree {
	if depth > maxDepth {
		depth = maxDepth
	}
	if depth < 1 {
		depth = 1
	}
	t := tree{
		depth:            depth,
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
		if len(ss) != 2 {
			log.Fatal(errors.New("error input"))
		}

		t.addItem(ss[0], ss[1])
	}
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
		item = &treeItem{module: t.modInfo.GetModuleAddIfEmpty(name)}
		t.items[name] = item
	}

	if t.rootItem == nil {
		t.rootItem = item
	}

	if len(childName) > 0 {
		item.children = append(item.children, t.addItem(childName, ""))
	}

	return item
}
