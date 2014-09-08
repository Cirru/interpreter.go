
package interpreter

func (env *scope) _bool(xs sequence) (ret unitype) {
  ret.Type = uniBool
  ret.Value = false
  if token, ok := xs[0].(token); ok {
    trueValues := []string{"true", "yes", "right", "1"}
    for _, text := range trueValues {
      if text == token.Text {
        ret.Value = true
      }
    }
    return
  } else {
    panic("failed to parse bool")
  }
  return
}
