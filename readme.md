### Getting started
Go version: 1.18

```
go get -u github.com/tlsproxy/filterx
```

generate structs
```go
type Custom struct {
    Name    string
    Age     int
}

gen := filterx.NewGeneratorOptions(
    WithOutput("./output"),
    WithGormHelper(),
    WithPackages([]string{"other package"}),
)

err := gen.Generate(Custom{})

if err != nil {
    panic(err)
}
```
