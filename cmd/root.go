package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vc60er/deptree/internal/moduleinfo"
	"github.com/vc60er/deptree/internal/tree"
)

// Execute starts the main code
// - get all upgradable modules: "go list -u -m -json all" and only those with newer version
// - filter list of "go mod graph" (all children with parent) to all upradeable children
// - print all parents needs to upgrade for usage of its (direct) upgradable children
func Execute() {
	// TODO:
	// - colored output
	// - check go version
	showAll := flag.Bool("a", false, "show all dependencies, not only with upgrade")
	colored := flag.Bool("c", false, "upgrade candidates will be marked yellow")
	depth := flag.Int("d", 3, "max depth of dependencies")
	visualizeTrimmed := flag.Bool("t", false, "visualize trimmed tree by '└─...'")
	printJSON := flag.Bool("json", false, "print JSON instead of tree")
	graphFile := flag.String("graph", "", "path to file created e.g. by 'go mod graph > grapphfile.txt'")
	upgradeFile := flag.String("upgrade", "", "path to file created e.g. by 'go list -u -m -json all > upgradefile.txt'")
	flag.Parse()

	info := moduleinfo.NewInfo()
	info.Fill(getUpgradeContent(*upgradeFile))

	tree := tree.NewTree(*depth, *visualizeTrimmed, *showAll, *colored, *info)
	file := getGraphFile(*graphFile)
	defer file.Close()
	tree.Fill(file)

	info.Adjust()

	tree.Print(*printJSON)
}

// getUpgradeContent gets the JSON content from go list call or upgrade file
func getUpgradeContent(upgradeFile string) []byte {
	// get all modules including upgrade versions
	var goListCallJSONContent []byte
	if len(upgradeFile) == 0 {
		var outbuf, errbuf bytes.Buffer
		cmd := exec.Command("go", "list", "-u", "-m", "-json", "all")
		cmd.Stdout = &outbuf
		cmd.Stderr = &errbuf
		if err := cmd.Run(); err != nil {
			log.Fatal(fmt.Sprintf("%v, %s", err, errbuf.String()))
		}
		goListCallJSONContent = outbuf.Bytes()
	} else {
		var err error
		if upgradeFile, err = filepath.Abs(upgradeFile); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("use upgrade file %s\n", upgradeFile)
		if goListCallJSONContent, err = ioutil.ReadFile(upgradeFile); err != nil {
			log.Fatal(err)
		}
	}

	return goListCallJSONContent
}

// getGraphFile gets the file handle to access content from STDIN or graph file
func getGraphFile(graphFile string) (file *os.File) {
	var err error
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
	return
}
