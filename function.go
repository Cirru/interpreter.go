
package interpreter

func (env *scope) fn(xs sequence) (ret unitype) {
  ret.Type = uniFn
  args, ok := xs[0].(sequence)
  if !ok {
    panic("function expects args in sequence")
  }
  ret.Value = context{env, args, xs[1:]}
  return
}

func (env *scope) call(xs sequence) (ret unitype) {
  fn := env.getValue(xs[0])
  if fn.Type != uniFn {
    panic("calling a non-function")
  }
  item, _ := fn.Value.(context)
  runtime := &scope{}
  for i, arg := range item.args {
    tok, _ := arg.(token)
    (*runtime)[uni(tok.Text)] = env.getValue(xs[i+1])
  }
  for _, line := range item.code {
    if exp, ok := line.(sequence); ok {
      ret = Evaluate(runtime, exp)
    } else {
      panic("calling not sequence")
    }
  }
  return
}
