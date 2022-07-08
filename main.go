package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pkg struct {
	name   string
	childs []string
}

type PkgTree struct {
	mapPkg map[string]*Pkg
	root   string
	depth  int
}

func NewPkgTree(depth int) *PkgTree {
	return &PkgTree{mapPkg: make(map[string]*Pkg), depth: depth}
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
		childPath = strings.Replace(childPath, "┌", "|", -1)

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

var pDepth = flag.Int("d", 3, "depth of dependence")

func main() {
	flag.Parse()

	graphFile := ""
	if flag.NArg() == 0 {
	} else if flag.NArg() == 1 {
		graphFile = flag.Arg(0)
	} else {
		flag.Usage()
		os.Exit(1)
	}

	var err error
	var file *os.File
	if len(graphFile) == 0 {
		file = os.Stdin
	} else {
		file, err = os.Open(graphFile)
		if err != nil {
			log.Fatal(err)
		}
	}
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

	fmt.Println("package:", root.name)
	fmt.Println("dependence tree:\n")
	childLen := len(root.childs)
	for i, c := range root.childs {
		head := "├── "
		if i == 0 {
			head = "┌── "
		} else if i == childLen-1 {
			head = "└── "
		}
		pkgTree.printTree(head, c)
	}
}
