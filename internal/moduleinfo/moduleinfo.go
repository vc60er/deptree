package moduleinfo

import (
	"encoding/json"
	"log"
	"strings"
)

// Info stores all attributes for retrieve module information
type Info struct {
	modules map[string]*Module
}

// NewInfo creates a new instance
func NewInfo() *Info {
	i := Info{
		modules: make(map[string]*Module),
	}
	return &i
}

// Fill adds given content to the modules map.
func (i *Info) Fill(goListCallJSONContent []byte) {
	// add missing paranthesis
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

	return
}

// GetModuleAddIfEmpty returns a module if exist, otherwise add it to the map before return
func (i *Info) GetModuleAddIfEmpty(name string) *Module {
	if m, ok := i.modules[name]; ok {
		return m
	}
	splitName := strings.Split(name, "@")
	m := Module{Path: splitName[0]}
	if len(splitName) == 2 {
		m.Version = splitName[1]
	}
	i.modules[name] = &m
	return &m
}

// Adjust modules after all is parsed
func (i *Info) Adjust() {
	for _, mod1 := range i.modules {
		for _, mod2 := range i.modules {
			if mod1.Path == mod2.Path && mod1.Version != mod2.Version {
				mod1.related = append(mod1.related, mod2)
			}
		}
	}
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
