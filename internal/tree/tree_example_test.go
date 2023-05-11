package tree

import (
	"io"
	"strings"

	"github.com/vc60er/deptree/internal/moduleinfo"
	"github.com/vc60er/deptree/internal/verbose"
)

func ExamplePrint_allTrimmed() {
	const (
		showDroppedChild = false
		visualizeTrimmed = true
		showAll          = true
		colored          = false
	)
	v := verbose.Verbose{}
	i := moduleinfo.NewInfo(v)
	t := NewTree(v, 2, showDroppedChild, visualizeTrimmed, showAll, colored, *i)
	t.Fill(graphStringReader())
	t.Print(false)
	// Output:
	// dependency tree with depth 2 for package: github.com/vc60er/deptree, least 2 trimmed item(s)
	// * tree of duplicate children not drawn (-f not set)
	//
	// github.com/vc60er/deptree
	//  └── github.com/stretchr/testify@v1.8.2
	//       ├── github.com/davecgh/go-spew@v1.1.1
	//       ├── github.com/pmezard/go-difflib@v1.0.0
	//       ├── github.com/stretchr/objx@v0.5.0
	//       │    └── 1 more ...
	//       └── gopkg.in/yaml.v3@v3.0.1
	//            └── 1 more ...
}

func ExamplePrint_allTrimmedSomeMore() {
	const (
		showDroppedChild = false
		visualizeTrimmed = true
		showAll          = true
		colored          = false
	)
	v := verbose.Verbose{}
	i := moduleinfo.NewInfo(v)
	t := NewTree(v, 1, showDroppedChild, visualizeTrimmed, showAll, colored, *i)
	t.Fill(graphStringReader())
	t.Print(false)
	// Output:
	// dependency tree with depth 1 for package: github.com/vc60er/deptree, least 4 trimmed item(s)
	// * tree of duplicate children not drawn (-f not set)
	//
	// github.com/vc60er/deptree
	//  └── github.com/stretchr/testify@v1.8.2
	//       └── 4 more ...
}

func ExamplePrint_all() {
	const (
		showDroppedChild = false
		visualizeTrimmed = false
		showAll          = true
		colored          = false
	)
	v := verbose.Verbose{}
	i := moduleinfo.NewInfo(v)
	t := NewTree(v, 2, showDroppedChild, visualizeTrimmed, showAll, colored, *i)
	t.Fill(graphStringReader())
	t.Print(false)
	// Output:
	// dependency tree with depth 2 for package: github.com/vc60er/deptree
	// * no visualization for trimmed tree (-t not set)
	// * tree of duplicate children not drawn (-f not set)
	//
	// github.com/vc60er/deptree
	//  └── github.com/stretchr/testify@v1.8.2
	//       ├── github.com/davecgh/go-spew@v1.1.1
	//       ├── github.com/pmezard/go-difflib@v1.0.0
	//       ├── github.com/stretchr/objx@v0.5.0
	//       └── gopkg.in/yaml.v3@v3.0.1
}

func ExamplePrint() {
	const (
		showDroppedChild = false
		visualizeTrimmed = false
		showAll          = false
		colored          = false
	)
	v := verbose.Verbose{}
	i := moduleinfo.NewInfo(v)
	i.Fill(upgradeContent())

	t := NewTree(v, 3, showDroppedChild, visualizeTrimmed, showAll, colored, *i)
	t.Fill(graphStringReader())

	i.Adjust()

	t.Print(false)
	// Output:
	// dependency tree with depth 3 for package: github.com/vc60er/deptree, least 2 trimmed item(s)
	// * no visualization for trimmed tree (-t not set)
	// * only upgradable items with parents are shown (-a not set)
	// * duplicate children not shown (-a not set)
	//
	// github.com/vc60er/deptree (go1.15)
	//  └── github.com/stretchr/testify@v1.8.2 (go1.13)
	//       ├── github.com/stretchr/objx@v0.5.0 (go1.12)
	//       │    └── github.com/stretchr/testify@v1.8.0 => [v1.8.2] (go1.13)
	//       └── gopkg.in/yaml.v3@v3.0.1
	//            └── gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405 => [v1.0.0-20201130134442-10cb98267c6c]
}

func upgradeContent() []byte {
	jsonContent := []string{
		`{"Path": "github.com/vc60er/deptree","GoVersion": "1.15"}`,
		`{"Path": "github.com/davecgh/go-spew",	"Version": "v1.1.1"}`,
		`{"Path": "github.com/pmezard/go-difflib","Version": "v1.0.0"}`,
		`{"Path": "github.com/stretchr/objx","Version": "v0.5.0",	"GoVersion": "1.12"}`,
		`{"Path": "github.com/stretchr/testify","Version": "v1.8.2","GoVersion": "1.13"}`,
		`{"Path": "gopkg.in/check.v1","Version": "v0.0.0-20161208181325-20d25e280405","Update": {"Path": "gopkg.in/check.v1","Version": "v1.0.0-20201130134442-10cb98267c6c","Time": "2020-11-30T13:44:42Z"}}`,
		`{"Path": "gopkg.in/yaml.v3","Version": "v3.0.1"}`,
	}
	return []byte(strings.Join(jsonContent, "\n"))
}

func graphStringReader() io.Reader {
	content := []string{
		"github.com/vc60er/deptree github.com/stretchr/testify@v1.8.2",
		"github.com/stretchr/testify@v1.8.2 github.com/davecgh/go-spew@v1.1.1",
		"github.com/stretchr/testify@v1.8.2 github.com/pmezard/go-difflib@v1.0.0",
		"github.com/stretchr/testify@v1.8.2 github.com/stretchr/objx@v0.5.0",
		"github.com/stretchr/testify@v1.8.2 gopkg.in/yaml.v3@v3.0.1",
		"github.com/stretchr/objx@v0.5.0 github.com/stretchr/testify@v1.8.0",
		"gopkg.in/yaml.v3@v3.0.1 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405",
	}
	return strings.NewReader(strings.Join(content, "\n"))
}
