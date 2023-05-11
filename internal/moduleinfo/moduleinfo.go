package moduleinfo

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/vc60er/deptree/internal/verbose"
)

// Info stores all attributes for retrieve module information
type Info struct {
	verbose verbose.Verbose
	modules map[string]*Module
}

// NewInfo creates a new instance
func NewInfo(verbose verbose.Verbose) *Info {
	i := Info{
		verbose: verbose,
		modules: make(map[string]*Module),
	}
	return &i
}

// Fill adds given content to the modules map.
func (i *Info) Fill(goListCallJSONContent []byte) {
	// add missing parenthesis
	goListCallJSONContent = append(append([]byte{'['}, goListCallJSONContent...), ']')
	// fix missing commas
	goListCallJSONContent = []byte(strings.ReplaceAll(string(goListCallJSONContent), "}\n{", "},{"))

	var moduleList []Module
	if err := json.Unmarshal(goListCallJSONContent, &moduleList); err != nil {
		log.Fatal(err)
	}

	for idx, module := range moduleList {
		i.modules[module.Name()] = &moduleList[idx]
	}
	i.verbose.Log1f("%d modules collected", len(i.modules))
}

// GetModuleAddIfEmpty returns a module if exist, otherwise add it to the map before return
func (i *Info) GetModuleAddIfEmpty(name string) *Module {
	if m, ok := i.modules[name]; ok {
		return m
	}
	splitName := strings.Split(name, "@")
	m := Module{Path: splitName[0]}
	const lenWithVersion = 2
	if len(splitName) == lenWithVersion {
		m.Version = splitName[1]
	}
	i.modules[name] = &m
	return &m
}

// Adjust modules after all is parsed
func (i *Info) Adjust() {
	var cnt int
	for _, mod1 := range i.modules {
		for _, mod2 := range i.modules {
			if mod1.Path == mod2.Path && mod1.Version != mod2.Version {
				cnt++
				mod1.related = append(mod1.related, mod2)
			}
		}
	}
	i.verbose.Log1f("%d adjustments done", cnt)
	i.verbose.Log1f("%d modules after adjust", len(i.modules))
}

// Print the module of the given path for debugging purposes
func (i *Info) Print(path string) {
	log.Println("len modules", len(i.modules))
	for _, m := range i.modules {
		if m.Path == path {
			log.Println(m)
		}
	}
}
