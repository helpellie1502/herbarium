package project

import (
	"os"
	"path/filepath"
)

func Scan(root string, ignoredDirs []string, ignoredFiles []string, ignoredExts []string) (*Project, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	newProject := &Project{}
	newProject.name = filepath.Base(absRoot)
	newProject.root = absRoot

	if err := walk(newProject, absRoot, ignoredDirs, ignoredFiles, ignoredExts); err != nil {
		return nil, err
	}
	return newProject, nil
}

func walk(p *Project, dir string, ignoredDirs []string, ignoredFiles []string, ignoredExts []string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			relPath, err := filepath.Rel(p.root, filepath.Join(dir, entry.Name()))
			if err != nil {
				return err
			}

			ignored := contains(ignoredDirs, entry.Name())

			p.tree = append(p.tree, Entry{
				path:      relPath,
				isDir:     true,
				isIgnored: ignored,
			})
			if ignored {
				continue
			}
			if err := walk(p, filepath.Join(dir, entry.Name()), ignoredDirs, ignoredFiles, ignoredExts); err != nil {
				return err
			}
		} else {
			relPath, err := filepath.Rel(p.root, filepath.Join(dir, entry.Name()))
			if err != nil {
				return err
			}

			ignored := contains(ignoredFiles, entry.Name()) || contains(ignoredExts, filepath.Ext(entry.Name()))

			p.tree = append(p.tree, Entry{
				path:      relPath,
				isDir:     false,
				isIgnored: ignored,
			})
		}
	}

	return nil
}

func contains(list []string, element string) bool {
	for _, listElement := range list {
		if listElement == element {
			return true
		}
	}
	return false
}
