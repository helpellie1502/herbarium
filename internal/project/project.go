package project

import (
	"errors"
	"fmt"
	"path/filepath"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type Entry struct {
	path      string
	isDir     bool
	isIgnored bool
}

type Project struct {
	root string
	name string
	tree []Entry
}

func (p *Project) Root() string {
	return p.root
}

func (p *Project) Name() string {
	return p.name
}

func (p *Project) Tree() []Entry {
	return p.tree
}

// TODO: удалить, временный метод для отладки
func (p *Project) TEntry(index int, sybDir string, sybFile string, projectNamePath bool) (string, error) {
	if index < 0 || index >= len(p.tree) {
		return "", ErrIndexOutOfRange
	}
	entry := p.tree[index]
	var entryType string
	var entryPath string

	if entry.isDir {
		entryType = sybDir
	} else {
		entryType = sybFile
	}

	if projectNamePath {
		entryPath = filepath.Join(p.name, entry.path)
	} else {
		entryPath = entry.path
	}

	result := fmt.Sprintf("%s %s", entryType, entryPath)
	return result, nil
}

func (e *Entry) Path() string {
	return e.path
}

func (e *Entry) IsDir() bool {
	return e.isDir
}

func (e *Entry) IsIgnored() bool {
	return e.isIgnored
}
