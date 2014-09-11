
package interpreter

func (env *scope) _if(xs sequence) unitype {
  cond := env.getValue(xs[0])
  if cond.Type != uniBool {
    panic("not a bool in condition")
  }
  isTrue, _ := cond.Value.(bool)
  var pointer int
  if isTrue {
    pointer = 1
  } else {
    pointer = 2
  }
  if pointer >= len(xs) {
    return uni(nil)
  }
  return env.getValue(xs[pointer])
}

func (env *scope) block(xs sequence) (ret unitype) {
  for _, line := range xs {
    ret = env.getValue(line)
  }
  return
}