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
	Childs []string `json:"Childs"`
}

type PkgTree struct {
	MapPkg map[string]*Pkg `json:"packages"`
	root   string
	depth  int
}

func NewPkgTree(depth int) *PkgTree {
	return &PkgTree{MapPkg: make(map[string]*Pkg), depth: depth}
}

func (p *PkgTree) Add(name string, child string) {
	pkg, ok := p.MapPkg[name]
	if !ok {
		pkg = &Pkg{name: name}
		p.MapPkg[name] = pkg
	}

	pkg.Childs = append(pkg.Childs, child)

	if len(p.root) == 0 {
		p.root = name
	}
}

func (p *PkgTree) GetPkg(name string) *Pkg {
	return p.MapPkg[name]
}

func (p *PkgTree) GetRootPkg() *Pkg {
	if len(p.root) == 0 {
		return nil
	}
	return p.MapPkg[p.root]
}

func (p *PkgTree) printTree(path string, name string) int {
	if len(path) > (p.depth+2)*5-1 {
		path = strings.Replace(path, "├", "└", -1)
		fmt.Printf("%s...\n", path)
		return -1
	}
	fmt.Printf("%s%s\n", path, name)

	child, ok := p.MapPkg[name]
	if ok && len(child.Childs) > 0 {
		childPath := path
		childPath = strings.Replace(childPath, "└", " ", -1)
		childPath = strings.Replace(childPath, "├", "│", -1)
		childPath = strings.Replace(childPath, "─", " ", -1)
		childPath = strings.Replace(childPath, "┌", "│", -1)

		childLen := len(child.Childs)
		for i, name := range child.Childs {
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
var printJSON = flag.Bool("json", false, "print JSON instead of tree")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [go-mod-graph-output-file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")

		flag.PrintDefaults()
	}

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

	if *printJSON {
		// Print dependencies in JSON format
		jsonTree, err := pkgTree.ToJSON()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", jsonTree)
	} else {
		// Print dependencies in Tree format
		fmt.Println("package:", root.name)
		fmt.Println("dependence tree:\n")
		childLen := len(root.Childs)
		for i, c := range root.Childs {
			head := "├── "
			if i == 0 {
				head = "┌── "
			} else if i == childLen-1 {
				head = "└── "
			}
			pkgTree.printTree(head, c)
		}
	}

}
