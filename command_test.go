package cirruGopher

import "testing"

func TestCommand(t *testing.T) {
	err := Interpret()
	if err != nil {
		t.Errorf("Runtime error", err)
	}
}
