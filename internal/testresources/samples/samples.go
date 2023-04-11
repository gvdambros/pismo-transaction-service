package samples

import (
	"path/filepath"
	"runtime"
)

var (
	_, p, _, _ = runtime.Caller(0)

	// Path root folder for samples
	Path = filepath.Dir(p)
)
