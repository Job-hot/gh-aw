//go:build !integration

package httpresponsebodyclose_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/github/gh-aw/pkg/linters/httpresponsebodyclose"
)

func TestHTTPResponseBodyClose(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, httpresponsebodyclose.Analyzer, "httpresponsebodyclose")
}
