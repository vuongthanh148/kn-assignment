// Package testhelper provides shared functionality and constants to be used in
// Discovery tests. It should only be imported by test files.
package testhelper

import (
	"path/filepath"
	"runtime"
)

// TestDataPath returns a path corresponding to a path relative to the calling
// test file. For convenience, rel is assumed to be "/"-delimited.
//
// It panics on failure.
func TestDataPath(rel string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to determine relative path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(filename), filepath.FromSlash(rel)))
}
