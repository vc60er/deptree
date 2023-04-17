package tree

import (
	"encoding/json"
	"fmt"
	"log"
)

// note: try to avoid the word "package" at Go level for variables etc.
type jsonTree struct {
	Modules []jsonModule `json:"packages"`
}

type jsonModule struct {
	Name     string   `json:"name"`
	Children []string `json:"children"`
}

// printJSON shows the given tree items as json
func (t *tree) printJSON() {
	jt := jsonTree{}
	jt.createAndAddModules(t)
	jsonContent, err := json.MarshalIndent(jt, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", jsonContent)
}

// createAndAddModules creates a object list for marshal to json
func (jt *jsonTree) createAndAddModules(t *tree) {
	for name, cs := range t.items {
		if len(cs.children) == 0 {
			continue
		}
		m := jsonModule{Name: name}
		for _, c := range cs.children {
			m.Children = append(m.Children, c.module.Name())
		}
		jt.Modules = append(jt.Modules, m)
	}
}
