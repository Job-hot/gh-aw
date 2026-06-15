// Package fprintstderrinlibrary implements a Go analysis linter that flags
// fmt.Fprint-style calls writing to os.Stderr in library (pkg/) packages.
package fprintstderrinlibrary

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

	"github.com/github/gh-aw/pkg/linters/internal/astutil"
	"github.com/github/gh-aw/pkg/linters/internal/filecheck"
	"github.com/github/gh-aw/pkg/linters/internal/nolint"
)

var Analyzer = &analysis.Analyzer{
	Name:     "fprintstderrinlibrary",
	Doc:      "reports fmt.Fprintf/Fprintln/Fprint calls to os.Stderr in library packages where pkg/logger should be used instead",
	URL:      "https://github.com/github/gh-aw/tree/main/pkg/linters/fprintstderrinlibrary",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var fprintFuncs = map[string]bool{
	"Fprint": true, "Fprintf": true, "Fprintln": true,
}

func run(pass *analysis.Pass) (any, error) {
	pkgPath := pass.Pkg.Path()
	if strings.HasSuffix(pkgPath, "/main") || strings.Contains(pkgPath, "/cmd/") {
		return nil, nil
	}

	insp, err := astutil.Inspector(pass)
	if err != nil {
		return nil, err
	}
	noLintLinesByFile := nolint.BuildLineIndex(pass, "fprintstderrinlibrary")

	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}
		if filecheck.IsTestFile(pass.Fset.Position(call.Pos()).Filename) {
			return
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		ident, ok := sel.X.(*ast.Ident)
		if !ok {
			return
		}
		if ident.Name != "fmt" || !fprintFuncs[sel.Sel.Name] || len(call.Args) == 0 {
			return
		}
		argSel, ok := call.Args[0].(*ast.SelectorExpr)
		if !ok {
			return
		}
		argIdent, ok := argSel.X.(*ast.Ident)
		if !ok {
			return
		}
		if argIdent.Name == "os" && argSel.Sel.Name == "Stderr" {
			position := pass.Fset.PositionFor(call.Pos(), false)
			if nolint.HasDirective(position, noLintLinesByFile) {
				return
			}
			pass.ReportRangef(call, "fmt.%s(os.Stderr, ...) called in library package %s; use pkg/logger instead", sel.Sel.Name, pkgPath)
		}
	})

	return nil, nil
}
