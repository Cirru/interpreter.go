
package interpreter

// import "fmt"

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
  ctx, _ := fn.Value.(context)
  runtime := newScope(ctx.closure)
  for i, arg := range ctx.args {
    tok, _ := arg.(token)
    (*runtime.closure)[tok.Text] = env.getValue(xs[i+1])
  }
  ret = uni(nil)
  for _, line := range ctx.code {
    ret = runtime.getValue(line)
  }
  return
}
