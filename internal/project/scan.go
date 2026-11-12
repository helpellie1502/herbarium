package project

import (
	"os"
	"path/filepath"
)

func Scan(root string, ignoredDirs []string, ignoredFiles []string) (*Project, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	newProject := &Project{}
	newProject.name = filepath.Base(absRoot)
	newProject.root = absRoot

	if err := walk(newProject, absRoot, ignoredDirs, ignoredFiles); err != nil {
		return nil, err
	}
	return newProject, nil
}

func walk(p *Project, dir string, ignoredDirs []string, ignoredFiles []string) error {
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
			p.tree = append(p.tree, Entry{
				path:      relPath,
				isDir:     true,
				isIgnored: isIgnoredDir(ignoredDirs, entry.Name()),
			})
			if isIgnoredDir(ignoredDirs, entry.Name()) {
				continue
			}
			if err := walk(p, filepath.Join(dir, entry.Name()), ignoredDirs, ignoredFiles); err != nil {
				return err
			}
		} else {
			relPath, err := filepath.Rel(p.root, filepath.Join(dir, entry.Name()))
			if err != nil {
				return err
			}
			p.tree = append(p.tree, Entry{
				path:      relPath,
				isDir:     false,
				isIgnored: isIgnoredFile(ignoredFiles, entry.Name()),
			})
		}
	}

	return nil
}

func isIgnoredDir(ignoredDirs []string, dir string) bool {
	for _, ignoredDir := range ignoredDirs {
		if ignoredDir == dir {
			return true
		}
	}
	return false
}

func isIgnoredFile(ignoredFiles []string, file string) bool {
	for _, ignoredFile := range ignoredFiles {
		if ignoredFile == file {
			return true
		}
	}
	return false
}
