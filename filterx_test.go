package filterx

import "testing"

type Custom struct {
	Name string
	Age  int
}

func TestNewGeneratorOptions(t *testing.T) {
	g := NewGeneratorOptions(
		WithOutput("./output"),
		WithGormHelper(),
		WithPackages([]string{"other package"}),
	)
	err := g.Generate(Custom{})
	if err != nil {
		t.Fail()
	}
}
