package main

import (
	"github.com/budougumi0617/layer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(layer.Analyzer) }
