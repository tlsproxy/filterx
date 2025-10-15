package filterx

import (
	"testing"
)

type B struct {
	Field3 int `json:"field3" filterx:"field3"`
}

type A struct {
	Field2 int `json:"field2" filterx:"field2"`
	*B
}

type Custom struct {
	Field1 int `json:"field1" filterx:"field1"`
	A
}

type Custom1 struct {
	Field1 int `json:"field1" filterx:"field1"`
}

func TestNewGeneratorOptions(t *testing.T) {
	g := NewGeneratorOptions(
		WithOutput("./output"),
		WithGormHelper(),
		WithPackages([]string{"time"}),
	)
	err := g.Generate(Custom{}, Custom1{})
	if err != nil {
		t.Fail()
	}
}
