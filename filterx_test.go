package filterx

import (
	"testing"
)

type App struct {
	ID uint `json:"id" filterx:"id"`
}

type Custom struct {
	App
	Name string `json:"name" filterx:"name"`
	Age  int    `json:"age" filterx:"age"`
}

func TestNewGeneratorOptions(t *testing.T) {
	g := NewGeneratorOptions(
		WithOutput("./output"),
		WithGormHelper(),
		//WithPackages([]string{""}),
	)
	err := g.Generate(Custom{})
	if err != nil {
		t.Fail()
	}
}
