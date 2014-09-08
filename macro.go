
package interpreter

func (env *scope) macro(xs sequence) (ret unitype) {
  ret.Type = uniMacro
  if args, ok := xs[0].(sequence); ok {
    ret.Value = context{env, args, xs[1:]}
  }
  return
}

func (env *scope) expand(xs sequence) (ret unitype) {
  macro := env.get(xs[0:1])
  if macro.Type == uniMacro {
    if ctx, ok := macro.Value.(context); ok {
      runtime := scope{}
      runtime[unitype{uniString, "outer"}] = unitype{uniTable, env}
      for i, args := range ctx.args {
        if t, ok := args.(token); ok {
          if para, ok := xs[i+1].(token); ok {
            runtime[uni(t.Text)] = unitype{uniString, para.Text}
          } else {
            panic("should not be expression in args")
          }
        }
      }
      for _, line := range ctx.code {
        if exp, ok := line.([]interface{}); ok {
          ret = Evaluate(&runtime, exp)
        }
      }
      return
    }
  } else {
    panic("not macro")
  }
  return
}