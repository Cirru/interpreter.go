
package interpreter

import "github.com/Cirru/parser"

func (env *Env) macro(xs []interface{}) (ret unitype) {
  ret.Type = uniMacro
  if args, ok := xs[0].([]interface{}); ok {
    ret.Value = context{env, args, xs[1:]}
  }
  return
}

func (env *Env) expand(xs []interface{}) (ret unitype) {
  macro := env.get(xs[0:1])
  if macro.Type == uniMacro {
    if ctx, ok := macro.Value.(context); ok {
      runtime := Env{}
      runtime[unitype{uniString, "outer"}] = unitype{uniTable, env}
      for i, args := range ctx.args {
        if token, ok := args.(parser.Token); ok {
          if para, ok := xs[i+1].(parser.Token); ok {
            runtime[uni(token.Text)] = unitype{uniString, para.Text}
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