package moduleinfo

import (
	"fmt"
	"strconv"
	"strings"
)

type version struct {
	major    int
	minor    int
	patch    int
	extra    int
	extraStr string
}

func parseVersion(vStr string) (*version, error) {
	// v1.27.1
	// v0.0.0-20180917221912-90fa682c2a6e
	// v1.26.0-rc.1
	// v4.3.0+incompatible
	var err error
	v := version{}
	vSplit := strings.Split(strings.ReplaceAll(vStr, "v", ""), ".")
	if len(vSplit) < 3 {
		return nil, fmt.Errorf("version not parseable from '%s'", vStr)
	}
	if v.major, err = strconv.Atoi(vSplit[0]); err != nil {
		return nil, err
	}
	if v.minor, err = strconv.Atoi(vSplit[1]); err != nil {
		return nil, err
	}
	if v.patch, err = strconv.Atoi(vSplit[2]); err == nil {
		// the plain case "v1.27.1"
		return &v, nil
	}
	// check for "-" separators
	vSplitExt := strings.Split(vSplit[2], "-")
	if len(vSplitExt) > 1 {
		if v.patch, err = strconv.Atoi(vSplitExt[0]); err != nil {
			return nil, err
		}
		if v.extra, err = strconv.Atoi(vSplitExt[1]); err != nil {
			v.extra = 0
			v.extraStr = strings.Join(vSplitExt[1:], "-") // contains now "rc.1"
		} else {
			v.extraStr = strings.Join(vSplitExt[2:], "-") // contains now "90fa682c2a6e"
		}
	}
	// check for "+" separator
	vSplitExt = strings.Split(vSplit[2], "+")
	if len(vSplitExt) > 1 {
		if v.patch, err = strconv.Atoi(vSplitExt[0]); err != nil {
			return nil, err
		}
		if v.extra, err = strconv.Atoi(vSplitExt[1]); err != nil {
			v.extra = 0
			v.extraStr = strings.Join(vSplitExt[1:], "+") // contains now "incompatible"
		} else {
			v.extraStr = strings.Join(vSplitExt[2:], "+") // contains all remaining
		}
	}
	return &v, nil
}

func higherOrEqualVersion(vam, vbm *Module) *Module {
	vaStr := vam.Version
	vbStr := vbm.Version
	if vaStr == vbStr {
		return vam
	}
	va, err := parseVersion(vaStr)
	if err != nil {
		return vbm
	}
	vb, err := parseVersion(vbStr)
	if err != nil {
		return vam
	}
	if va.major > vb.major {
		return vam
	}
	if va.major < vb.major {
		return vbm
	}
	if va.minor > vb.minor {
		return vam
	}
	if va.minor < vb.minor {
		return vbm
	}
	if va.patch > vb.patch {
		return vam
	}
	if va.patch < vb.patch {
		return vbm
	}
	// all is equal, so compare the extra string
	if va.extra > vb.extra {
		return vam
	}
	if va.extra < vb.extra {
		return vbm
	}
	if va.extraStr > vb.extraStr {
		return vam
	}
	return vbm
}
