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
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFileAtomic writes file to temp and atomically move when everything else succeeds.
func WriteFileAtomic(file string, data []byte, perm os.FileMode) error {
	dir, name := filepath.Split(file)
	f, err := ioutil.TempFile(dir, name)
	if err != nil {
		return err
	}
	if _, err = f.Write(data); err != nil {
		return err
	}
	if err = f.Sync(); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	if err = os.Chmod(f.Name(), perm); err != nil {
		return err
	}
	if err = os.Rename(f.Name(), file); err != nil {
		return err
	}

	_ = os.Remove(f.Name())
	return nil
}

func AllFileLines(file string) ([]string, error) {
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return AllLines(bytes.NewReader(bts)), nil
}

func AllLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	var a []string
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	return a
}

func WriteUnixTextFile(filename string, data []byte, perm os.FileMode) error {
	x := bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
	return ioutil.WriteFile(filename, x, perm)
}
