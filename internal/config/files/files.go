package files

import (
	"path/filepath"
	"runtime"
)

var (
	_, p, _, _ = runtime.Caller(0)

	// Path root folder for config files
	Path = filepath.Dir(p)
)
