// Package httpresponsebodyclose implements a Go analysis linter that flags
// HTTP response bodies that are closed manually instead of via defer.
package httpresponsebodyclose

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

	"github.com/github/gh-aw/pkg/linters/internal/astutil"
	"github.com/github/gh-aw/pkg/linters/internal/filecheck"
)

// Analyzer is the http-response-body-close analysis pass.
var Analyzer = &analysis.Analyzer{
	Name:     "httpresponsebodyclose",
	Doc:      "reports HTTP response Body.Close() calls that are not deferred, which can cause resource leaks on early return or panic",
	URL:      "https://github.com/github/gh-aw/tree/main/pkg/linters/httpresponsebodyclose",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	insp, err := astutil.Inspector(pass)
	if err != nil {
		return nil, err
	}

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		inspectHTTPResponseFuncDecl(pass, n)
	})

	return nil, nil
}

type responseVarState struct {
	assignPos      token.Pos
	hasDefer       bool
	hasManualClose bool
}

func inspectHTTPResponseFuncDecl(pass *analysis.Pass, n ast.Node) {
	fn, ok := n.(*ast.FuncDecl)
	if !ok || fn.Body == nil {
		return
	}

	pos := pass.Fset.PositionFor(fn.Pos(), false)
	if filecheck.IsTestFile(pos.Filename) {
		return
	}

	responseVars := make(map[types.Object]*responseVarState)

	ast.Inspect(fn.Body, func(node ast.Node) bool {
		return analyzeResponseNode(pass, responseVars, node)
	})

	for _, state := range responseVars {
		if state.hasManualClose && !state.hasDefer {
			pass.Report(analysis.Diagnostic{
				Pos:     state.assignPos,
				Message: "HTTP response Body.Close() should be deferred immediately after error check to prevent resource leaks",
			})
		}
	}
}

func analyzeResponseNode(pass *analysis.Pass, responseVars map[types.Object]*responseVarState, node ast.Node) bool {
	if node == nil {
		return false
	}

	// Do not descend into function literals — closures are independent execution
	// contexts and should be analyzed separately to avoid false positives.
	if _, ok := node.(*ast.FuncLit); ok {
		return false
	}

	if assign, ok := node.(*ast.AssignStmt); ok {
		trackResponseAssignment(pass, responseVars, assign)
	}

	if deferStmt, ok := node.(*ast.DeferStmt); ok {
		if obj := getHTTPBodyCloseObject(pass, deferStmt.Call); obj != nil {
			if state, found := responseVars[obj]; found {
				state.hasDefer = true
			}
		}
	}

	// Non-deferred resp.Body.Close() in expression statements.
	if exprStmt, ok := node.(*ast.ExprStmt); ok {
		if call, ok := exprStmt.X.(*ast.CallExpr); ok {
			markManualBodyClose(pass, responseVars, call)
		}
	}

	// Non-deferred resp.Body.Close() captured in an assignment (e.g. closeErr := resp.Body.Close()).
	if assign, ok := node.(*ast.AssignStmt); ok {
		for _, rhs := range assign.Rhs {
			if call, ok := rhs.(*ast.CallExpr); ok {
				markManualBodyClose(pass, responseVars, call)
			}
		}
	}

	return true
}

func trackResponseAssignment(pass *analysis.Pass, responseVars map[types.Object]*responseVarState, assign *ast.AssignStmt) {
	// Track any binding whose LHS has type *net/http.Response. Unlike
	// fileclosenotdeferred (which tracks explicit os.Open/os.Create call sites),
	// this intentionally keys on the assigned variable's type so responses
	// returned by helper functions are also covered.
	for _, lhs := range assign.Lhs {
		ident, ok := lhs.(*ast.Ident)
		if !ok || ident.Name == "_" {
			continue
		}
		obj := pass.TypesInfo.ObjectOf(ident)
		if obj == nil {
			continue
		}
		if !isHTTPResponseType(obj.Type()) {
			continue
		}
		if prev, exists := responseVars[obj]; exists && prev.hasManualClose && !prev.hasDefer {
			pass.Report(analysis.Diagnostic{
				Pos:     prev.assignPos,
				Message: "HTTP response Body.Close() should be deferred immediately after error check to prevent resource leaks",
			})
		}
		assignPos := ident.Pos()
		if len(assign.Rhs) == 1 {
			if call, ok := assign.Rhs[0].(*ast.CallExpr); ok {
				assignPos = call.Pos()
			}
		}
		responseVars[obj] = &responseVarState{assignPos: assignPos}
	}
}

func markManualBodyClose(pass *analysis.Pass, responseVars map[types.Object]*responseVarState, call *ast.CallExpr) {
	obj := getHTTPBodyCloseObject(pass, call)
	if obj == nil {
		return
	}
	if state, found := responseVars[obj]; found {
		state.hasManualClose = true
	}
}

// getHTTPBodyCloseObject returns the types.Object for the *http.Response variable
// in a resp.Body.Close() call, or nil if the call does not match that pattern.
// Known limitation: body aliasing (body := resp.Body; body.Close()) is not
// detected because the selector chain no longer starts from the *http.Response
// variable directly.
func getHTTPBodyCloseObject(pass *analysis.Pass, call *ast.CallExpr) types.Object {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Close" {
		return nil
	}
	// The receiver must be an expression of the form resp.Body.
	bodyExpr, ok := sel.X.(*ast.SelectorExpr)
	if !ok || bodyExpr.Sel.Name != "Body" {
		return nil
	}
	ident, ok := bodyExpr.X.(*ast.Ident)
	if !ok {
		return nil
	}
	obj := pass.TypesInfo.ObjectOf(ident)
	if obj == nil {
		return nil
	}
	if !isHTTPResponseType(obj.Type()) {
		return nil
	}
	return obj
}

// isHTTPResponseType reports whether t is *net/http.Response.
func isHTTPResponseType(t types.Type) bool {
	ptr, ok := t.(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}
	obj := named.Obj()
	return obj.Name() == "Response" && obj.Pkg() != nil && obj.Pkg().Path() == "net/http"
}
