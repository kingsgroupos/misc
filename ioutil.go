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
