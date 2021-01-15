// BSD 3-Clause License
//
// Copyright (c) 2020, Kingsgroup
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package misc

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ResetWorkingDirectory() error {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(exePath)
	if strings.HasSuffix(dir, fmt.Sprintf("%cbin", os.PathSeparator)) {
		dir = filepath.Dir(dir)
	} else if strings.HasSuffix(filepath.Dir(dir), fmt.Sprintf("%cbin", os.PathSeparator)) {
		dir = filepath.Dir(filepath.Dir(dir))
	}
	return os.Chdir(dir)
}

func AbsDir(p string, nthR2L int) (string, error) {
	d, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	for i := 0; i < nthR2L; i++ {
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
