package moduleinfo

import (
	"fmt"
	"time"
)

// Module is module information, see "go help list"
type Module struct {
	Path      string       `json:"Path"`
	Version   string       `json:"Version"`
	Versions  []string     `json:"Versions"`  // available module versions (with -versions)
	Replace   *Module      `json:"Replace"`   // replaced by this module
	Time      *time.Time   `json:"Time"`      // time version was created
	Update    *Module      `json:"Update"`    // available update, if any (with -u)
	Main      bool         `json:"Main"`      // is this the main module?
	Indirect  bool         `json:"Indirect"`  // is this module only an indirect dependency of main module?
	Dir       string       `json:"Dir"`       // directory holding files for this module, if any
	GoMod     string       `json:"GoMod"`     // path to go.mod file used when loading this module, if any
	GoVersion string       `json:"GoVersion"` // go version used in module, e.g. "1.11"
	Error     *ModuleError `json:"Error"`     // error loading module
	related   []*Module    // contains all found modules with the same Path but different version, this means all found versions
}

// ModuleError is the content of module loading error
type ModuleError struct {
	Err string `json:"Err"` // the error itself
}

// GetUpdateModule returns a possible update candidate and nil if there was no higher version found or there is a
// parsing error.
func (m *Module) GetUpdateModule() *Module {
	latestModule := m.Update
	if latestModule == nil {
		latestModule = m
	}
	for _, modRelated := range m.related {
		latestModule = higherOrEqualVersion(modRelated, latestModule)
		if modRelated.Update != nil {
			latestModule = higherOrEqualVersion(modRelated.Update, latestModule)
		}
	}

	if m.Version == latestModule.Version {
		return nil
	}

	return latestModule
}

// Name gets the module name from path and version
func (m *Module) Name() string {
	if m.Version != "" {
		return fmt.Sprintf("%s@%s", m.Path, m.Version)
	}
	return m.Path
}
