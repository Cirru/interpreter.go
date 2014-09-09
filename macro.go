
package interpreter

// import "fmt"

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
      runtime[uni("outer")] = uni(env)
      for i, arg := range ctx.args {
        if t, ok := arg.(token); ok {
          if para, ok := xs[i+1].(token); ok {
            runtime[uni(t.Text)] = uni(para.Text)
          } else {
            panic("should not be expression in args")
          }
        }
      }
      for _, line := range ctx.code {
        if seq, ok := line.(sequence); ok {
          // fmt.Println(seq)
          ret = Evaluate(&runtime, seq)
        }
      }
      return
    }
  } else {
    panic("not macro")
  }
  return
}