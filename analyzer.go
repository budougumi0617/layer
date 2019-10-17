package layer

import (
	"encoding/json"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

/**
 * define layer graph by JSON array.
 * For example, the sample graph of "The Clean Architecture" is below:
 * https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
 * [
 *   "external_interfaces",
 *   "web",
 *   "devices",
 *   "db",
 *   "ui",
 *   [
 *     "controllers",
 *     "gateways",
 *     "presenters",
 *     [
 *       "use_cases",
 *       [
 *         "entity"
 *       ]
 *     ]
 *   ]
 * ]
 */
var jsonString = `[ "external","db","ui", [ "controllers", [ "usecases", [ "entity" ] ] ] ]` // -jsonlayer flag

func init() {
	Analyzer.Flags.StringVar(&jsonString, "jsonlayer", jsonString, "jsonlayer defines layer hierarchy by JSON array")
}

// Analyzer confirms whether the packages follow to the layer structure.
var Analyzer = &analysis.Analyzer{
	Name: "layer",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// Doc defines analysis messages. this message is shown by "layer help" command.
const Doc = "layer checks whether there are dependencies that illegal cross-border the layer structure. The layer structure is defined as a JSON array using the -jsonlayer option."

func run(pass *analysis.Pass) (interface{}, error) {
	l := &Layer{}
	if err := json.Unmarshal([]byte(jsonString), l); err != nil {
		return nil, err
	}
	currentPackage := pass.Pkg.Path()
found:
	for ; l.Inside != nil; l = l.Inside {
		for _, ln := range l.Packages {
			if strings.Contains(currentPackage, ln) {
				break found
			}
		}
	}

	if l.Inside == nil || l.Inside.Inside == nil {
		return nil, nil
	}

	il := l.Inside.Inside
	for _, f := range pass.Files {
		if strings.HasSuffix(pass.Fset.File(f.Pos()).Name(), "_test.go") {
			continue
		}
		for _, i := range f.Imports {
			// TODO: ignore standard packages.
			// memo: https://github.com/golang/go/blob/6cba4dbf80012c272cb04bd878dfba251d9bb05c/src/cmd/go/internal/modload/build.go#L30
			if invalid(il, i.Path.Value) {
				pass.Reportf(i.Pos(), "%s must not include %s", currentPackage, i.Path.Value)
			}
		}
	}

	return nil, nil
}

func invalid(l *Layer, s string) bool {
	path := strings.Trim(s, "\"")
	for {
		if include(l, path) {
			return true
		}
		if l.Inside == nil {
			return false
		}
		l = l.Inside
	}
}

func include(l *Layer, name string) bool {
	for _, p := range l.Packages {
		if strings.Contains(name, p) {
			return true
		}
	}
	return false
}
