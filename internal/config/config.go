package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	IgnoredDirs  []string `json:"ignored_dirs"`
	IgnoredFiles []string `json:"ignored_files"`
	IgnoredExts  []string `json:"ignored_exts"`

	OutputDir string `json:"output_dir"`
}

func Default() *Config {
	ctf := &Config{
		IgnoredDirs:  []string{".git"},
		IgnoredFiles: []string{".env"},
		IgnoredExts:  []string{".png"},

		OutputDir: ".",
	}

	return ctf
}

func Path() (string, error) {
	ctfDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	ctfPath := filepath.Join(ctfDir, "herbarium", "config.json")
	return ctfPath, nil
}
