package layer_test

import (
	"path/filepath"
	"testing"

	"github.com/budougumi0617/layer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestOption(t *testing.T) {
	if err := layer.Analyzer.Flags.Set("jsonlayer", `["handler", ["usecase", ["domain", "repository"]]]`); err != nil {
		t.Fatal(err)
	}
	testdata := filepath.Join(analysistest.TestData(), "option")
	analysistest.Run(t, testdata, layer.Analyzer, "handler")
}

func TestNested(t *testing.T) {
	if err := layer.Analyzer.Flags.Set("jsonlayer", `["handler", ["usecase", ["domain", "repository"]]]`); err != nil {
		t.Fatal(err)
	}
	testdata := filepath.Join(analysistest.TestData(), "nested")
	analysistest.Run(t, testdata, layer.Analyzer, "handler/...")
}

func TestImportPath(t *testing.T) {
	if err := layer.Analyzer.Flags.Set("jsonlayer", `["a", ["b", ["c"]]]`); err != nil {
		t.Fatal(err)
	}
	testdata := filepath.Join(analysistest.TestData(), "importpath")
	analysistest.Run(t, testdata, layer.Analyzer, "a/...")
}
