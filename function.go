
package interpreter

func (env *scope) fn(xs sequence) (ret unitype) {
  ret.Type = uniFn
  if args, ok := xs[0].(sequence); ok {
    ret.Value = context{env, args, xs[1:]}
  }
  return
}

func (env *scope) call(xs sequence) (ret unitype) {
  fn := env.get(xs[0:1])
  if fn.Type == uniFn {
    if item, ok := fn.Value.(context); ok {
      runtime := scope{}
      for i, para := range item.args {
        // println("i is:", i)
        if token, ok := para.(token); ok {
          runtime[uni(token.Text)] = env.get(xs[i+1:i+2])
        }
      }
      for _, line := range item.code {
        if exp, ok := line.(sequence); ok {
          ret = Evaluate(&runtime, exp)
        } else {
          panic("calling not sequence")
        }
      }
      return
    }
  } else {
    panic("not function")
  }
  return
}
