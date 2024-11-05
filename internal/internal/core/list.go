package core

import (
	"os"
	"path/filepath"

	"github.com/goexl/gfx/internal/internal/core/internal"
	"github.com/goexl/gfx/internal/internal/kernel"
	"github.com/goexl/gfx/internal/internal/param"
)

type List struct {
	params *param.List
	file   *internal.File
}

func NewList(params *param.List) *List {
	return &List{
		params: params,
		file:   internal.NewFile(params.File),
	}
}

func (l *List) All() (files []string) {
	files = make([]string, 0, 8)
	for _, directories := range l.params.Directories {
		directory := filepath.Join(directories...)
		patterns := l.file.Patterns(directory)
		for _, pattern := range patterns {
			if all, ge := filepath.Glob(pattern); nil == ge && 0 != len(all) && kernel.FileTypeAll == l.params.Type {
				files = append(files, all...)
			} else if nil == ge && 0 != len(all) {
				files = append(files, l.check(all)...)
			}
		}
	}

	return
}

func (l *List) check(files []string) (checked []string) {
	checked = make([]string, 0, len(files))
	for _, file := range files {
		if info, se := os.Stat(file); nil != se {
			continue
		} else if kernel.FileTypeDirectory == l.params.Type && info.IsDir() {
			checked = append(checked, file)
		} else if kernel.FileTypeFile == l.params.Type {
			checked = append(checked, file)
		}
	}

	return
}
