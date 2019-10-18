# layer
[![godoc](https://godoc.org/github.com/budougumi0617/layer?status.svg)][godoc]
[![github actions reviewdog badges](https://github.com/budougumi0617/layer/workflows/reviewdog/badge.svg)][actions_reviwdog]
[![github actions test badges](https://github.com/budougumi0617/layer/workflows/test/badge.svg)][actions_test]
[![GolangCI](https://golangci.com/badges/github.com/budougumi0617/layer.svg)][golangci]

[godoc]:https://godoc.org/github.com/budougumi0617/layer
[actions_reviwdog]:https://github.com/budougumi0617/layer/actions?workflow=reviewdog
[actions_test]:https://github.com/budougumi0617/layer/actions?workflow=test
[golangci]:https://golangci.com/r/github.com/budougumi0617/layer

`layer` checks whether there are dependencies that illegal cross-border the layer structure. The layer structure is defined as a JSON array using the -jsonlayer option.

## Install

You can get `layer` by `go get` command.

```bash
$ go get -u github.com/budougumi0617/layer
```

## QuickStart

`layer` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which layer) -layer.jsonlayer "[\"webhandler\", [\"usecase\", [\"domain\"]]]" github.com/path/to/target/pacakge
```

When Go is lower than 1.12, just run `layer` command with the package name (import path) and option.

```bash
$ layer -jsonlayer "[\"webhandler\", [\"usecase\", [\"domain\"]]]" github.com/path/to/target/pacakge
```

## Analyzer

`layer` checks whether there are dependencies that illegal cross-border the layer structure. The layer structure is defined as a JSON array using the -jsonlayer option.

Typical layered architectures(ex: the clean architecture, the onion architecture, etc) have the dependency rules that source code dependencies can only point inwards. Go can enforce this rule, because Go does not allow circular dependencies.
However, other rule cannot enforce by Go language specification, it rule is that every element in a layer must depends only the elements on the same layer or on the “just beneath” layer.
The `layer` command checks whether there are dependencies that illegal cross-border the layer structure.

For example, the sample layer graph of "The Clean Architecture" is shown below JSON array:
https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
```json
[
  "externalinterfaces",
  "web",
  "devices",
  "db",
  "ui",
  [
    "controllers",
    "gateways",
    "presenters",
    [
      "usecases",
      [
        "entity"
      ]
    ]
  ]
]
```

If there are below illegal codes:

```go
package web

import (
	"encoding/json"
	"net/http"

	"usecases" // "web" package must not import the next lower layer(ex: controllers).
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int
	}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	repo := repository.UserRepository()
	if err := repo.Delete(req.ID); err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
```

`layer` shows below error.
```bash
$ go vet -vettool=$(which layer) -layer.jsonlayer "[\"externalinterfaces\", \"web\", \"devices\", \"db\", \"ui\", [ \"controllers\", \"gateways\", \"presenters\", [ \"usecases\", [ \"entity\" ] ] ] ]" ./...
# github.com/budougumi0617/clean_architecture/web
web/delete_handler.go:7:2: github.com/budougumi0617/clean_architecture/web must not import "github.com/budougumi0617/clean_architecture/usecases"
```

## Description
```bash
$ layer help layer
layer: layer checks whether there are dependencies that illegal cross-border the layer structure. The layer structure is defined as a JSON array using the -jsonlayer option.

Analyzer flags:

  -layer.jsonlayer string
        jsonlayer defines layer hierarchy by JSON array (default "[ \"external\",\"db\",\"ui\", [ \"controllers\", [ \"usecases\", [ \"entity\" ] ] ] ]")
```

## TODO
- [ ] Support to load JSON array from file

## Contribution
1. Fork ([https://github.com/budougumi0617/layer/fork](https://github.com/budougumi0617/layer/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/budougumi0617/layer/blob/master/LICENSE)

## Author
[Yoichiro Shimizu(@budougumi0617)](https://github.com/budougumi0617)

