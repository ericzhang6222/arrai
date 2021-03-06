package test

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/arr-ai/arrai/pkg/ctxfs"
	"github.com/arr-ai/arrai/rel"
	"github.com/arr-ai/arrai/syntax"
)

// Test runs all tests in the subtree of path and returns the results.
func Test(ctx context.Context, w io.Writer, path string) (Results, error) {
	results := Results{}
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "_test.arrai") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return Results{}, err
	}

	if len(files) == 0 {
		return results, fmt.Errorf("no tests found in %v", path)
	}

	fmt.Fprintf(w, "Tests:\n%s\n", strings.Join(files, "\n"))

	for _, file := range files {
		bytes, err := ctxfs.ReadFile(ctxfs.SourceFsFrom(ctx), file)
		if err != nil {
			return results, err
		}
		result, err := syntax.EvaluateExpr(ctx, file, string(bytes))
		if err != nil {
			fmt.Fprintf(w, "\nfailed test: %s\n", err)
			results.Add(Result{file: file, pass: false})
		} else {
			results.Add(Result{file: file, pass: isRecursivelyTrue(result)})
		}
	}

	return results, nil
}

// isRecursivelyTrue returns true if every leaf value of val is true (not just truthy).
func isRecursivelyTrue(val rel.Value) bool {
	switch v := val.(type) {
	case rel.GenericSet:
		if !v.IsTrue() {
			return false
		}
		if v.Count() == 1 && v.Has(rel.NewTuple()) {
			return true
		}
		for _, item := range v.OrderedValues() {
			if !isRecursivelyTrue(item) {
				return false
			}
		}
		return true
	case rel.Array:
		if v.Count() == 0 {
			return false
		}
		for _, item := range v.Values() {
			if !isRecursivelyTrue(item) {
				return false
			}
		}
		return true
	case rel.Dict:
		if v.Count() == 0 {
			return false
		}
		for _, entry := range v.OrderedEntries() {
			if !isRecursivelyTrue(entry.MustGet(rel.DictValueAttr)) {
				return false
			}
		}
		return true
	case rel.Tuple:
		for e := v.Enumerator(); e.MoveNext(); {
			_, attr := e.Current()
			if !isRecursivelyTrue(attr) {
				return false
			}
		}
		return true
	}
	return false
}
