package main

import (
	"encoding/json"
)

func (p *PkgTree) ToJSON() (treeJson []byte, err error) {
	treeJson, err = json.MarshalIndent(p, "", "    ")
	return
}
