package misc

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ResetWorkingDirectory() error {
	exePath := os.Args[0]
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		return err
	}

	dir := filepath.Dir(absPath)
	if strings.HasSuffix(dir, fmt.Sprintf("%cbin", os.PathSeparator)) {
		dir = filepath.Dir(dir)
	} else if strings.HasSuffix(filepath.Dir(dir), fmt.Sprintf("%cbin", os.PathSeparator)) {
		dir = filepath.Dir(filepath.Dir(dir))
	}
	return os.Chdir(dir)
}

func AbsDir(p string, n int) (string, error) {
	d, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	for i := 0; i < n; i++ {
		d = filepath.Dir(d)
	}
	return d, nil
}

func AllFiles(patterns []string, suffix string, singleDir bool) ([]string, error) {
	var files []string
	var unique = make(map[string]struct{})
	for _, p := range patterns {
		matches, err := filepath.Glob(p)
		if err != nil {
			return nil, err
		}
		for _, f := range matches {
			abs, err := filepath.Abs(f)
			if err != nil {
				return nil, err
			}
			cleaned := filepath.Clean(abs)
			if _, ok := unique[cleaned]; ok {
				continue
			}
			if stat, err := os.Stat(f); err != nil {
				return nil, err
			} else if stat.IsDir() {
				continue
			} else if suffix != "" && !strings.HasSuffix(f, suffix) {
				continue
			}
			files = append(files, f)
			unique[cleaned] = struct{}{}
		}
	}

	if singleDir {
		for _, f := range files {
			if filepath.Dir(f) != filepath.Dir(files[0]) {
				return nil, fmt.Errorf("all files must reside in a single directory. f1: %s, f2: %s", files[0], f)
			}
		}
	}

	sort.Strings(files)
	return files, nil
}

func FindFile(name string) error {
	stat, err := os.Stat(name)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("%s is a directory", name)
	}
	return nil
}

func FindDirectory(name string) error {
	stat, err := os.Stat(name)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("%s is a NOT a directory", name)
	}
	return nil
}
