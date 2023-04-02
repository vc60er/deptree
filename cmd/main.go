package main

import (
	"flag"

	"github.com/vc60er/deptree/internal/tree"
)

func main() {
	// TODO:
	// - get all upgradable modules: "go list -u -m -json all" and only those with newer version
	// - filter list of "go mod graph" (all children with parent) to all upradeable children
	// - print all parents needs to upgrade for usage of its (direct) upgradable children
	// - colored output
	// - check go version
	depth := flag.Int("d", 3, "max depth of dependencies")
	visualizeTrimmed := flag.Bool("t", false, "visualize trimmed tree by '└─...'")
	printJSON := flag.Bool("json", false, "print JSON instead of tree")
	graphFile := flag.String("graph", "", "path to file created e.g. by 'go mod graph > grapphfile.txt'")
	flag.Parse()

	tree := tree.NewTree(*depth, *visualizeTrimmed)
	tree.Fill(*graphFile)
	tree.Print(*printJSON)
}
