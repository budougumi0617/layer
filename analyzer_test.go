package layer_test

import (
	"testing"

	"github.com/budougumi0617/layer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, layer.Analyzer, "a")
}
