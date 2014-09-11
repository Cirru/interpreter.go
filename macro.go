
package interpreter

// import "fmt"

func (env *scope) macro(xs sequence) (ret unitype) {
  ret.Type = uniMacro
  args, ok := xs[0].(sequence)
  if !ok {
   panic("macro excepts args in sequence")
  }
  code := xs[1:]
  ret.Value = context{env, args, code}
  return
}

func (env *scope) expand(xs sequence) (ret unitype) {
  macro := env.getValue(xs[0])
  if macro.Type != uniMacro {
    panic("not macro")
  }
  ctx, _ := macro.Value.(context)
  runtime := &scope{}
  (*runtime)[uni("outer")] = uni(env)
  for i, arg := range ctx.args {
    key := env.getKey(arg)
    tok, ok := xs[i].(token)
    if !ok {
      panic("expand excepts token arguments")
    }
    (*runtime)[key] = uni(tok.Text)
  }
  for _, line := range ctx.code {
    seq, _ := line.(sequence)
    ret = Evaluate(runtime, seq)
  }
  return
}