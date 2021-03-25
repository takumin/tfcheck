package fswalk

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

func FsWalk(dir string, excludes []string) ([]string, error) {
	var dirs []string

	if err := filepath.WalkDir(dir, func(name string, info fs.DirEntry, err error) error {
		if info.IsDir() {
			return nil
		}

		for _, v := range excludes {
			if strings.Index(name, v) > 0 {
				return nil
			}
		}

		if path.Ext(name) != ".tf" {
			return nil
		}

		dirs = append(dirs, name)

		return nil
	}); err != nil {
		return nil, fmt.Errorf("Failed to filepath.WalkDir(): %w", err)
	}

	return dirs, nil
}
