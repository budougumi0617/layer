package layer_test

import (
	"path/filepath"
	"testing"

	"github.com/budougumi0617/layer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestOption(t *testing.T) {
	testdata := filepath.Join(analysistest.TestData(), "option")
	layer.Analyzer.Flags.Set("jsonlayer", `["handler", ["usecase", ["domain", "repository"]]]`)
	analysistest.Run(t, testdata, layer.Analyzer, "handler")
}

func TestNested(t *testing.T) {
	layer.Analyzer.Flags.Set("jsonlayer", `["handler", ["usecase", ["domain", "repository"]]]`)
	testdata := filepath.Join(analysistest.TestData(), "nested")
	analysistest.Run(t, testdata, layer.Analyzer, "handler/...")
}

func TestImportPath(t *testing.T) {
	layer.Analyzer.Flags.Set("jsonlayer", `["a", ["b", ["c"]]]`)
	testdata := filepath.Join(analysistest.TestData(), "importpath")
	analysistest.Run(t, testdata, layer.Analyzer, "a/...")
}
