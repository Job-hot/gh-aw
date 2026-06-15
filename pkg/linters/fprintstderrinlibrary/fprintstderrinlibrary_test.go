//go:build !integration

package fprintstderrinlibrary_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/github/gh-aw/pkg/linters/fprintstderrinlibrary"
)

func TestFprintStderrInLibrary(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, fprintstderrinlibrary.Analyzer, "fprintstderrinlibrary")
}
