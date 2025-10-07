
Range、EQ、NEQ、GT、GE、LT、LE、In、Like、NIn

### Getting started
Go version: 1.18
```
go get -u github.com/tlsproxy/filterx
```


```go
package main

import "github.com/tlsproxy/filterx"

type Custom struct {
	Name string `json:"name" filterx:"name"`
	Age  int    `json:"age" filterx:"age"`
}

func main() {
	gen := filterx.NewGeneratorOptions(
		filterx.WithOutput("./filter"),
		//filterx.WithGormHelper(),
		//filterx.WithPackages([]string{"other package"}),
	)
	err := gen.Generate(Custom{})
	if err != nil {
		panic(err)
	}
}
```



```go
package main

import (
	"encoding/json"
	"filter/filter"
	"github.com/tlsproxy/filterx"
)

func main() {
	custom := &filter.CustomFilterx{
		Name: filterx.EQ[string]{
			Value: "filterx",
		},
		Age: filterx.Range[int]{
			Max: 20,
			Min: 10,
		},
	}

	j, err := json.MarshalIndent(custom, "", "  ")
	if err != nil {
		panic(err)
	}

	println(string(j))
}

```
