
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
  runtime := &scope{}
  (*runtime)[uni("parent")] = uni(ctx.closure)
  (*runtime)[uni("outer")] = uni(env)
  for i, arg := range ctx.args {
    tok, _ := arg.(token)
    (*runtime)[uni(tok.Text)] = env.getValue(xs[i+1])
  }
  ret = uni(nil)
  for _, line := range ctx.code {
    ret = runtime.getValue(line)
  }
  return
}

func (env *scope) method(xs sequence) (ret unitype) {
  target := env.getValue(xs[0])
  if target.Type != uniTable {
    panic("method expects a table")
  }
  area, _ := target.Value.(*scope)
  item := area.getValue(xs[1])
  if item.Type != uniFn {
    return item
  }
  ctx, _ := item.Value.(context)
  runtime := &scope{}
  (*runtime)[uni("parent")] = uni(ctx.closure)
  (*runtime)[uni("outer")] = uni(env)
  (*runtime)[uni("this")] = uni(area)
  for i, arg := range ctx.args {
    tok, _ := arg.(token)
    (*runtime)[uni(tok.Text)] = env.getValue(xs[i+1])
  }
  ret = uni(nil)
  for _, line := range ctx.code {
    ret = runtime.getValue(line)
  }
  return
}