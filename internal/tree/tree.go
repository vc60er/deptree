package tree

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/vc60er/deptree/internal/moduleinfo"
)

const (
	maxDepth    = 20
	depthMarker = "..."
)

type treeItem struct {
	info     moduleinfo.Module
	children []*treeItem
}

type tree struct {
	depth            int
	visualizeTrimmed bool
	items            map[string]*treeItem
	rootItem         *treeItem
}

// NewTree creates a new instance for tree visualization
func NewTree(depth int, visualizeTrimmed bool) *tree {
	if depth > maxDepth {
		depth = maxDepth
	}
	t := tree{
		depth:            depth,
		visualizeTrimmed: visualizeTrimmed,
		items:            make(map[string]*treeItem),
	}
	return &t
}

// Fill the tree with content from STDIN or file
func (t *tree) Fill(graphFile string) {
	var err error
	var file *os.File
	if len(graphFile) == 0 {
		file = os.Stdin
	} else {
		if graphFile, err = filepath.Abs(graphFile); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("use graph file %s\n", graphFile)
		file, err = os.Open(graphFile)
		if err != nil {
			log.Fatal(err)
		}
	}
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

func (t *tree) Print(asJson bool) {
	if t.rootItem == nil {
		return
	}

	if asJson {
		t.printJSON()
		return
	}

	t.printRoute()
}

func (t *tree) addItem(name string, childName string) *treeItem {
	item, ok := t.items[name]
	if !ok {
		item = &treeItem{info: moduleinfo.NewModule(name)}
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
