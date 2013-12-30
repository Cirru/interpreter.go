
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
)

type context struct {
  env *Env
  args cirru.List
  code cirru.List
}

func cirruBlock(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "block"
  if args, ok := xs[0].(cirru.List); ok {
    ret.Value = context{env, args, xs[1:]}
  }
  return
}

func cirruCall(env *Env, xs cirru.List) (ret Object) {
  block := cirruGet(env, xs[0:1])
  // debugPrint("block is", block)
  if block.Tag == "block" {
    if item, ok := block.Value.(context); ok {
      runtime := Env{}
      for i, para := range item.args {
        // println("i is:", i)
        // debugPrint(xs)
        if token, ok := para.(cirru.Token); ok {
          runtime[token.Text] = cirruGet(env, xs[i+1:i+2])
        }
      }
      for _, line := range item.code {
        if exp, ok := line.(cirru.List); ok {
          ret = Evaluate(&runtime, exp)
        }
      }
      return
    }
  } else {
    stop("not block")
  }
  return
}