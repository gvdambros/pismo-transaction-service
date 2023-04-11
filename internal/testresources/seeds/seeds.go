package seeds

import (
	"path/filepath"
	"runtime"
)

var (
	_, p, _, _ = runtime.Caller(0)

	// Path root folder for db seeds
	Path = filepath.Dir(p)
)
