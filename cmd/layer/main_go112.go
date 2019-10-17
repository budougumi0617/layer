// +build go1.12

package main

import (
	"flag"

	"github.com/budougumi0617/layer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

// Analyzers returns analyzers of layer.
func analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		layer.Analyzer,
	}
}

func main() {
	unitchecker.Main(analyzers()...)
}

func init() {
	// go vet always adds this flag for certain packages in the standard library,
	// which causes "flag provided but not defined" errors when running with
	// custom vet tools.
	// So we just declare it here and swallow the flag.
	// See https://github.com/golang/go/issues/34053 for details.
	// TODO: Remove this once above issue is resolved.
	flag.String("unsafeptr", "", "")
}
