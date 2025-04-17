### Getting started
Go version: 1.18

```
go get -u github.com/tlsproxy/filterx
```

generate structs
```go
package main

import "github.com/tlsproxy/filterx"

type Custom struct {
	Name string
	Age  int
}

func main() {
	gen := filterx.NewGeneratorOptions(
		filterx.WithOutput("./filter"),
		filterx.WithGormHelper(),
		//filterx.WithPackages([]string{"other package"}),
	)
	err := gen.Generate(Custom{})
	if err != nil {
		panic(err)
	}
}
```

```go
go mod tidy
```
