
package interpreter

func (env *scope) _bool(xs sequence) (ret unitype) {
  ret.Type = uniBool
  tok, ok := xs[0].(token)
  if !ok {
    panic("failed to parse bool")
  }
  switch tok.Text {
  case "true", "yes":
    ret.Value = true
  default:
    ret.Value = false
  }
  return
}
