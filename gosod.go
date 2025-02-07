package gosod

import (
	"github.com/morganhein/gosod/internal/templatedir"
	"io/fs"
)

// New creates a new TemplateDir structure for the given filesystem
func New(fs fs.FS) *templatedir.TemplateDir {
	return templatedir.New(fs)
}
