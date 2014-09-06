
package interpreter

import "github.com/Cirru/parser"

type context struct {
  env *Env
  args []interface{}
  code []interface{}
}

func cirruFunction(env *Env, xs []interface{}) (ret Object) {
  ret.Tag = cirruTypeFunction
  if args, ok := xs[0].([]interface{}); ok {
    ret.Value = context{env, args, xs[1:]}
  }
  return
}

func cirruCall(env *Env, xs []interface{}) (ret Object) {
  function := cirruGet(env, xs[0:1])
  if function.Tag == cirruTypeFunction {
    if item, ok := function.Value.(context); ok {
      runtime := Env{}
      for i, para := range item.args {
        // println("i is:", i)
        if token, ok := para.(parser.Token); ok {
          runtime[token.Text] = cirruGet(env, xs[i+1:i+2])
        }
      }
      for _, line := range item.code {
        if exp, ok := line.([]interface{}); ok {
          ret = Evaluate(&runtime, exp)
        }
      }
      return
    }
  } else {
    stop("not function")
  }
  return
}