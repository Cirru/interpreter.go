
package cirruGopher

import "testing"

func TestSqrt(t *testing.T) {
  err := Interpret()
  if err != nil {
    t.Errorf("Runtime error", err)
  }
}