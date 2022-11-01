package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pkg struct {
	name   string   `json:"package_name"`
	childs []string `json:"package_childs"`
}

type PkgTree struct {
	mapPkg map[string]*Pkg `json:"packages"`
	root   string          `json:"-"`
	depth  int             `json:"-"`
}

func NewPkgTree(depth int) *PkgTree {
	return &PkgTree{mapPkg: make(map[string]*Pkg), depth: depth}
}

func (p *PkgTree) ToJSON() (treeJson []byte, err error) {
	treeJson, err = json.Marshal(p)
	return
}

func (p *PkgTree) Add(name string, child string) {
	pkg, ok := p.mapPkg[name]
	if !ok {
		pkg = &Pkg{name: name}
		p.mapPkg[name] = pkg
	}

	pkg.childs = append(pkg.childs, child)

	if len(p.root) == 0 {
		p.root = name
	}
}

func (p *PkgTree) GetPkg(name string) *Pkg {
	return p.mapPkg[name]
}

func (p *PkgTree) GetRootPkg() *Pkg {
	if len(p.root) == 0 {
		return nil
	}
	return p.mapPkg[p.root]
}

func (p *PkgTree) printTree(path string, name string) int {
	if len(path) > (p.depth+2)*5-1 {
		path = strings.Replace(path, "├", "└", -1)
		fmt.Printf("%s...\n", path)
		return -1
	}
	fmt.Printf("%s%s\n", path, name)

	child, ok := p.mapPkg[name]
	if ok && len(child.childs) > 0 {
		childPath := path
		childPath = strings.Replace(childPath, "└", " ", -1)
		childPath = strings.Replace(childPath, "├", "│", -1)
		childPath = strings.Replace(childPath, "─", " ", -1)
		childPath = strings.Replace(childPath, "┌", "│", -1)

		childLen := len(child.childs)
		for i, name := range child.childs {
			corner := ""
			if i == childLen-1 {
				corner = childPath + " └── "
			} else {
				corner = childPath + " ├── "
			}
			if p.printTree(corner, name) == -1 {
				break
			}
		}
	}

	return 0
}

var pDepth = flag.Int("d", 3, "max depth of dependence")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [output file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")

		flag.PrintDefaults()
	}

	flag.Parse()

	outputFile := ""
	if flag.NArg() == 0 {
	} else if flag.NArg() == 1 {
		outputFile = flag.Arg(0)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	var err error
	var file = os.Stdin
	reader := bufio.NewReader(file)
	pkgTree := NewPkgTree(*pDepth)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		ss := strings.Split(string(line), " ")
		if len(ss) != 2 {
			log.Fatal(errors.New("error input"))
		}

		pkgTree.Add(ss[0], ss[1])
	}

	root := pkgTree.GetRootPkg()
	if root == nil {
		return
	}

	// Export the JSON format of the tree.
	jsonTree, err := pkgTree.ToJSON()
	err = os.WriteFile("dep.json", jsonTree, 0644)
	if err != nil {
		panic(err)
	}
	//fmt.Println("package:", root.name)
	//fmt.Println("dependence tree:\n")
	//childLen := len(root.childs)
	//for i, c := range root.childs {
	//	head := "├── "
	//	if i == 0 {
	//		head = "┌── "
	//	} else if i == childLen-1 {
	//		head = "└── "
	//	}
	//	pkgTree.printTree(head, c)
	//}
}
